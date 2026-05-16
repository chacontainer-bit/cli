import {
  Package, RefreshCw, ShoppingCart, Leaf, AlertTriangle, FolderOpen,
  Droplets, Wrench, XCircle, Clock, CheckCircle, Truck, MapPin, Map,
  FileText, CheckSquare, Layers, AlertCircle, FlaskConical, Shield,
  ClipboardCheck, Recycle, TrendingUp, TrendingDown, Minus
} from 'lucide-react'
import {
  BarChart, Bar, XAxis, YAxis, Tooltip, ResponsiveContainer,
  PieChart, Pie, Cell, Legend
} from 'recharts'
import { useApp } from '../context/AppContext'
import { kpis, chartDataContainers, monthlyData, roles } from '../data/mockData'

const iconMap = {
  Package, RefreshCw, ShoppingCart, Leaf, AlertTriangle, FolderOpen,
  Droplets, Wrench, XCircle, Clock, CheckCircle, Truck, MapPin, Map,
  FileText, CheckSquare, Layers, AlertCircle, FlaskConical, Shield,
  ClipboardCheck, Recycle,
}

const alertColors = {
  critico: { border: 'border-red-500', bg: 'bg-red-500/10', dot: 'bg-red-500', label: 'CRÍTICO', text: 'text-red-400' },
  advertencia: { border: 'border-amber-500', bg: 'bg-amber-500/10', dot: 'bg-amber-500', label: 'AVISO', text: 'text-amber-400' },
  info: { border: 'border-blue-500', bg: 'bg-blue-500/10', dot: 'bg-blue-500', label: 'INFO', text: 'text-blue-400' },
}

const activityColors = {
  'QR Escaneado': 'text-emerald-400',
  'Pedido Creado': 'text-blue-400',
  'Lavado Completo': 'text-cyan-400',
  'NC Registrada': 'text-red-400',
  'Despacho': 'text-amber-400',
  'Alerta Crítica': 'text-red-400',
}

function KPICard({ label, value, delta, trend, icon }) {
  const Icon = iconMap[icon] || Package
  const isUp = trend === 'up'
  const isDown = trend === 'down'
  const TrendIcon = isUp ? TrendingUp : isDown ? TrendingDown : Minus
  const trendColor = isUp ? 'text-emerald-400' : isDown ? 'text-red-400' : 'text-gray-500'

  return (
    <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4 flex flex-col gap-3">
      <div className="flex items-start justify-between">
        <div className="w-9 h-9 bg-[#f97316]/10 rounded-lg flex items-center justify-center">
          <Icon size={18} className="text-[#f97316]" />
        </div>
        <div className={`flex items-center gap-1 text-xs font-mono ${trendColor}`}>
          <TrendIcon size={12} />
          <span>{delta}</span>
        </div>
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
  const { role, alerts, activity } = useApp()
  const roleData = roles.find(r => r.id === role)
  const roleKpis = kpis[role] || kpis.supply_chain

  const now = new Date()
  const hour = now.getHours()
  const greeting = hour < 12 ? 'Buenos días' : hour < 18 ? 'Buenas tardes' : 'Buenas noches'

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-xl font-bold text-[#f9fafb]">
            {greeting}, <span style={{ color: roleData?.color }}>{roleData?.name}</span>
          </h1>
          <p className="text-[#6b7280] text-sm mt-0.5">
            {now.toLocaleDateString('es-CL', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })}
            {' · '}
            <span className="font-mono">{now.toLocaleTimeString('es-CL', { hour: '2-digit', minute: '2-digit' })}</span>
          </p>
        </div>
        <div className="flex items-center gap-2 px-3 py-1.5 bg-emerald-500/10 border border-emerald-500/30 rounded-full">
          <div className="w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse" />
          <span className="text-emerald-400 text-xs font-medium">Sistema Operativo</span>
        </div>
      </div>

      {/* KPI Grid */}
      <div className="grid grid-cols-2 lg:grid-cols-3 gap-4">
        {roleKpis.map((kpi, i) => <KPICard key={i} {...kpi} />)}
      </div>

      {/* Middle row: Alerts + Activity */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
        {/* Alerts */}
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl">
          <div className="px-4 py-3 border-b border-[#1f2937] flex items-center justify-between">
            <h2 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2">
              <AlertTriangle size={15} className="text-[#f97316]" />
              Alertas Activas
            </h2>
            <span className="text-xs font-mono text-[#6b7280]">{alerts.length} total</span>
          </div>
          <div className="divide-y divide-[#1f2937] max-h-72 overflow-y-auto">
            {alerts.map(alert => {
              const cfg = alertColors[alert.type]
              return (
                <div key={alert.id} className={`flex gap-3 p-3 border-l-2 ${cfg.border} ${cfg.bg}`}>
                  <div className={`w-1.5 h-1.5 rounded-full mt-1.5 flex-shrink-0 ${cfg.dot}`} />
                  <div className="min-w-0 flex-1">
                    <div className="flex items-center gap-2 mb-0.5">
                      <span className={`text-xs font-bold font-mono ${cfg.text}`}>{cfg.label}</span>
                      <span className="text-[#6b7280] text-xs">{alert.module}</span>
                    </div>
                    <p className="text-[#d1d5db] text-xs leading-relaxed">{alert.message}</p>
                  </div>
                  <span className="text-[#6b7280] text-xs font-mono flex-shrink-0 pt-0.5">{alert.time}</span>
                </div>
              )
            })}
          </div>
        </div>

        {/* Activity Feed */}
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl">
          <div className="px-4 py-3 border-b border-[#1f2937] flex items-center justify-between">
            <h2 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2">
              <RefreshCw size={15} className="text-[#3b82f6]" />
              Actividad Reciente
            </h2>
            <div className="flex items-center gap-1.5">
              <div className="w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse" />
              <span className="text-xs text-emerald-400 font-mono">LIVE</span>
            </div>
          </div>
          <div className="divide-y divide-[#1f2937] max-h-72 overflow-y-auto">
            {activity.slice(0, 10).map((item, i) => (
              <div key={i} className="flex items-start gap-3 p-3">
                <span className="text-[#6b7280] text-xs font-mono flex-shrink-0 pt-0.5 w-10">{item.time}</span>
                <div className="min-w-0 flex-1">
                  <span className={`text-xs font-semibold ${activityColors[item.event] || 'text-[#9ca3af]'}`}>
                    {item.event}
                  </span>
                  <p className="text-[#d1d5db] text-xs mt-0.5 truncate">{item.detail}</p>
                </div>
                <span className="text-[#6b7280] text-xs flex-shrink-0">{item.user}</span>
              </div>
            ))}
          </div>
        </div>
      </div>

      {/* Charts row */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
        {/* Bar Chart */}
        <div className="lg:col-span-2 bg-[#111827] border border-[#1f2937] rounded-xl p-4">
          <h2 className="text-sm font-semibold text-[#f9fafb] mb-4">Ciclos de Lavado & Pedidos — Últimos 7 Meses</h2>
          <ResponsiveContainer width="100%" height={200}>
            <BarChart data={monthlyData} barGap={2}>
              <XAxis dataKey="mes" tick={{ fill: '#6b7280', fontSize: 11 }} axisLine={false} tickLine={false} />
              <YAxis tick={{ fill: '#6b7280', fontSize: 11 }} axisLine={false} tickLine={false} />
              <Tooltip content={<CustomTooltip />} cursor={{ fill: 'rgba(249,115,22,0.05)' }} />
              <Bar dataKey="ciclos" name="Ciclos" fill="#f97316" radius={[3, 3, 0, 0]} />
              <Bar dataKey="pedidos" name="Pedidos" fill="#3b82f6" radius={[3, 3, 0, 0]} />
            </BarChart>
          </ResponsiveContainer>
        </div>

        {/* Pie Chart */}
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
          <h2 className="text-sm font-semibold text-[#f9fafb] mb-4">Estado Flota Actual</h2>
          <ResponsiveContainer width="100%" height={200}>
            <PieChart>
              <Pie
                data={chartDataContainers}
                cx="50%"
                cy="50%"
                innerRadius={50}
                outerRadius={75}
                paddingAngle={2}
                dataKey="value"
              >
                {chartDataContainers.map((entry, index) => (
                  <Cell key={`cell-${index}`} fill={entry.fill} />
                ))}
              </Pie>
              <Tooltip
                formatter={(value, name) => [value, name]}
                contentStyle={{ backgroundColor: '#1a2235', border: '1px solid #1f2937', borderRadius: '8px', color: '#f9fafb', fontSize: '12px' }}
              />
            </PieChart>
          </ResponsiveContainer>
          <div className="space-y-1.5 mt-2">
            {chartDataContainers.slice(0, 4).map(d => (
              <div key={d.name} className="flex items-center justify-between text-xs">
                <div className="flex items-center gap-1.5">
                  <div className="w-2 h-2 rounded-full" style={{ backgroundColor: d.fill }} />
                  <span className="text-[#9ca3af]">{d.name}</span>
                </div>
                <span className="text-[#f9fafb] font-mono">{d.value}</span>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}
