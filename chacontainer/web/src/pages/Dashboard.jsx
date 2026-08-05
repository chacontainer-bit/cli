import { useEffect, useState } from 'react'
import {
  Package, Truck, AlertTriangle, ScanLine, Loader2, TrendingUp, Warehouse
} from 'lucide-react'
import {
  XAxis, YAxis, Tooltip, ResponsiveContainer,
  PieChart, Pie, Cell, LineChart, Line,
} from 'recharts'
import { api } from '../lib/api'
import { useAuth } from '../context/AuthContext'

const statusColors = {
  available: '#10b981',
  in_use: '#3b82f6',
  in_transit: '#f97316',
  maintenance: '#eab308',
  retired: '#6b7280',
  lost: '#ef4444',
}

const statusLabels = {
  available: 'Disponible',
  in_use: 'En uso',
  in_transit: 'En tránsito',
  maintenance: 'Mantenimiento',
  retired: 'Retirado',
  lost: 'Perdido',
}

function KPICard({ label, value, icon: Icon, color }) {
  return (
    <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4 flex flex-col gap-3">
      <div className="w-9 h-9 rounded-lg flex items-center justify-center" style={{ backgroundColor: `${color}1a` }}>
        <Icon size={18} style={{ color }} />
      </div>
      <div>
        <p className="text-2xl font-bold text-[#f9fafb] font-mono">{value}</p>
        <p className="text-[#6b7280] text-xs mt-0.5">{label}</p>
      </div>
    </div>
  )
}

const CustomTooltip = ({ active, payload, label }) => {
  if (active && payload && payload.length) {
    return (
      <div className="bg-[#1a2235] border border-[#1f2937] rounded-lg p-3">
        <p className="text-[#9ca3af] text-xs mb-1">{label}</p>
        {payload.map((p, i) => (
          <p key={i} className="text-sm font-mono" style={{ color: p.color }}>
            {p.name}: {p.value}
          </p>
        ))}
      </div>
    )
  }
  return null
}

export default function Dashboard() {
  const { user, tenant } = useAuth()
  const [summary, setSummary] = useState(null)
  const [assetTrend, setAssetTrend] = useState([])
  const [shipmentTrend, setShipmentTrend] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    let cancelled = false
    async function load() {
      setLoading(true)
      setError(null)
      try {
        const [s, at, st] = await Promise.all([
          api.dashboardSummary(),
          api.assetTrend(14),
          api.shipmentTrend(14),
        ])
        if (cancelled) return
        setSummary(s)
        setAssetTrend((at?.data || []).map(p => ({ date: p.date?.slice(5), activos: p.value })))
        setShipmentTrend((st?.data || []).map(p => ({ date: p.date?.slice(5), envios: p.value })))
      } catch (err) {
        if (!cancelled) setError(err.message)
      } finally {
        if (!cancelled) setLoading(false)
      }
    }
    load()
    return () => { cancelled = true }
  }, [])

  const now = new Date()
  const hour = now.getHours()
  const greeting = hour < 12 ? 'Buenos días' : hour < 18 ? 'Buenas tardes' : 'Buenas noches'

  const pieData = summary
    ? Object.entries(summary.assets_by_status).map(([status, value]) => ({
        name: statusLabels[status] || status, value, fill: statusColors[status] || '#6b7280',
      }))
    : []

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-xl font-bold text-[#f9fafb]">
            {greeting}, <span className="text-[#f97316]">{user?.name}</span>
          </h1>
          <p className="text-[#6b7280] text-sm mt-0.5">
            {tenant?.name} · {now.toLocaleDateString('es-MX', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })}
          </p>
        </div>
        <div className="flex items-center gap-2 px-3 py-1.5 bg-emerald-500/10 border border-emerald-500/30 rounded-full">
          <div className="w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse" />
          <span className="text-emerald-400 text-xs font-medium">Sistema Operativo</span>
        </div>
      </div>

      {error && (
        <div className="bg-red-500/10 border border-red-500/30 rounded-xl p-4 text-red-400 text-sm">
          No se pudo cargar el dashboard: {error}
        </div>
      )}

      {loading && !summary && (
        <div className="flex items-center gap-2 text-[#6b7280] text-sm py-12 justify-center">
          <Loader2 size={16} className="animate-spin" /> Cargando…
        </div>
      )}

      {summary && (
        <>
          <div className="grid grid-cols-2 lg:grid-cols-5 gap-4">
            <KPICard label="Activos totales" value={summary.total_assets} icon={Package} color="#f97316" />
            <KPICard label="En tránsito" value={summary.assets_in_transit} icon={Truck} color="#3b82f6" />
            <KPICard label="Envíos de hoy" value={summary.shipments_today} icon={TrendingUp} color="#10b981" />
            <KPICard label="Envíos atrasados" value={summary.late_shipments} icon={AlertTriangle} color="#ef4444" />
            <KPICard label="Escaneos 24h" value={summary.scan_events_24h} icon={ScanLine} color="#eab308" />
          </div>

          <div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
            <div className="lg:col-span-2 bg-[#111827] border border-[#1f2937] rounded-xl p-4">
              <h2 className="text-sm font-semibold text-[#f9fafb] mb-4">Altas de activos y envíos — últimos 14 días</h2>
              <ResponsiveContainer width="100%" height={200}>
                <LineChart data={mergeTrends(assetTrend, shipmentTrend)}>
                  <XAxis dataKey="date" tick={{ fill: '#6b7280', fontSize: 11 }} axisLine={false} tickLine={false} />
                  <YAxis tick={{ fill: '#6b7280', fontSize: 11 }} axisLine={false} tickLine={false} allowDecimals={false} />
                  <Tooltip content={<CustomTooltip />} />
                  <Line type="monotone" dataKey="activos" name="Activos" stroke="#f97316" strokeWidth={2} dot={false} />
                  <Line type="monotone" dataKey="envios" name="Envíos" stroke="#3b82f6" strokeWidth={2} dot={false} />
                </LineChart>
              </ResponsiveContainer>
            </div>

            <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
              <h2 className="text-sm font-semibold text-[#f9fafb] mb-4">Activos por estado</h2>
              {pieData.length === 0 ? (
                <p className="text-[#6b7280] text-xs text-center py-8">Sin activos aún</p>
              ) : (
                <>
                  <ResponsiveContainer width="100%" height={160}>
                    <PieChart>
                      <Pie data={pieData} cx="50%" cy="50%" innerRadius={45} outerRadius={70} paddingAngle={2} dataKey="value">
                        {pieData.map((entry, i) => <Cell key={i} fill={entry.fill} />)}
                      </Pie>
                      <Tooltip contentStyle={{ backgroundColor: '#1a2235', border: '1px solid #1f2937', borderRadius: '8px', color: '#f9fafb', fontSize: '12px' }} />
                    </PieChart>
                  </ResponsiveContainer>
                  <div className="space-y-1.5 mt-2">
                    {pieData.map(d => (
                      <div key={d.name} className="flex items-center justify-between text-xs">
                        <div className="flex items-center gap-1.5">
                          <div className="w-2 h-2 rounded-full" style={{ backgroundColor: d.fill }} />
                          <span className="text-[#9ca3af]">{d.name}</span>
                        </div>
                        <span className="text-[#f9fafb] font-mono">{d.value}</span>
                      </div>
                    ))}
                  </div>
                </>
              )}
            </div>
          </div>

          <div className="bg-[#111827] border border-[#1f2937] rounded-xl">
            <div className="px-4 py-3 border-b border-[#1f2937] flex items-center gap-2">
              <Warehouse size={15} className="text-[#f97316]" />
              <h2 className="text-sm font-semibold text-[#f9fafb]">Ocupación por planta</h2>
            </div>
            <div className="divide-y divide-[#1f2937]">
              {(summary.plant_occupancy || []).length === 0 && (
                <p className="text-[#6b7280] text-xs text-center py-6">Sin plantas registradas</p>
              )}
              {(summary.plant_occupancy || []).map(p => (
                <div key={p.plant_id} className="px-4 py-3 flex items-center justify-between gap-4">
                  <span className="text-[#d1d5db] text-sm">{p.plant_name}</span>
                  <div className="flex items-center gap-3 flex-1 max-w-xs">
                    <div className="h-2 bg-[#1f2937] rounded-full flex-1 overflow-hidden">
                      <div className="h-full bg-[#f97316] rounded-full" style={{ width: `${Math.min(100, p.occupancy_pct)}%` }} />
                    </div>
                    <span className="text-[#6b7280] text-xs font-mono w-20 text-right">{p.used_slots}/{p.total_slots}</span>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </>
      )}
    </div>
  )
}

function mergeTrends(assetTrend, shipmentTrend) {
  const byDate = {}
  for (const p of assetTrend) {
    byDate[p.date] = { date: p.date, activos: p.activos, envios: 0 }
  }
  for (const p of shipmentTrend) {
    byDate[p.date] = { ...(byDate[p.date] || { date: p.date, activos: 0 }), envios: p.envios }
  }
  return Object.values(byDate).sort((a, b) => a.date.localeCompare(b.date))
}
