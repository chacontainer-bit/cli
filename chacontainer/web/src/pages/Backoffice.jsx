import { Settings, CheckCircle, XCircle, Package, Wrench, Truck } from 'lucide-react'
import { useApp } from '../context/AppContext'
import { workshops, containers, products } from '../data/mockData'

function CapacityBar({ current, capacity }) {
  const pct = Math.round((current / capacity) * 100)
  const color = pct > 80 ? 'bg-red-500' : pct > 60 ? 'bg-amber-500' : 'bg-emerald-500'
  return (
    <div>
      <div className="flex justify-between text-xs mb-1">
        <span className="text-[#9ca3af] font-mono">{current}/{capacity}</span>
        <span className="text-[#6b7280]">{pct}%</span>
      </div>
      <div className="h-2 bg-[#1f2937] rounded-full overflow-hidden">
        <div className={`h-full rounded-full transition-all ${color}`} style={{ width: `${pct}%` }} />
      </div>
    </div>
  )
}

function QRFeed({ activity }) {
  const qrEvents = activity.filter(a =>
    a.event === 'QR Escaneado' || a.event === 'Lavado Completo' || a.event === 'Despacho'
  ).slice(0, 12)

  const eventColor = {
    'QR Escaneado': 'text-emerald-400',
    'Lavado Completo': 'text-cyan-400',
    'Despacho': 'text-amber-400',
  }

  return (
    <div className="bg-[#111827] border border-[#1f2937] rounded-xl">
      <div className="px-4 py-3 border-b border-[#1f2937] flex items-center justify-between">
        <h3 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2">
          <div className="w-2 h-2 bg-emerald-500 rounded-full animate-pulse" />
          Feed QR en Tiempo Real
        </h3>
        <span className="text-xs font-mono text-emerald-400">LIVE</span>
      </div>
      <div className="divide-y divide-[#1f2937] max-h-72 overflow-y-auto font-mono">
        {qrEvents.map((ev, i) => (
          <div key={i} className="px-4 py-2.5 flex items-center gap-3 text-xs">
            <span className="text-[#6b7280] w-10 flex-shrink-0">{ev.time}</span>
            <span className={`flex-shrink-0 w-24 ${eventColor[ev.event] || 'text-[#9ca3af]'}`}>{ev.event}</span>
            <span className="text-[#d1d5db] truncate flex-1">{ev.detail}</span>
            <span className="text-[#6b7280] flex-shrink-0">{ev.user}</span>
          </div>
        ))}
        {qrEvents.length === 0 && (
          <div className="py-8 text-center text-[#6b7280] text-sm">Esperando eventos...</div>
        )}
      </div>
    </div>
  )
}

export default function Backoffice() {
  const { activity } = useApp()

  const totalContainers = containers.length
  const byStatus = containers.reduce((acc, c) => {
    acc[c.status] = (acc[c.status] || 0) + 1
    return acc
  }, {})

  const inventorySummary = products.map(p => {
    const totalStock = Object.values(p.stock).reduce((a, b) => (b === 999 ? a : a + b), 0)
    const isLow = totalStock > 0 && totalStock < 20 && p.category !== 'Servicios'
    return { ...p, totalStock, isLow }
  })

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-xl font-bold text-[#f9fafb] flex items-center gap-2">
            <Settings size={20} className="text-[#f97316]" />
            Backoffice
          </h1>
          <p className="text-[#6b7280] text-sm">Vista operativa de talleres, inventario y actividad</p>
        </div>
      </div>

      {/* Summary KPIs */}
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
        {[
          { label: 'Contenedores Totales', value: totalContainers, icon: Package, color: '#f97316' },
          { label: 'Disponibles', value: byStatus.disponible || 0, icon: CheckCircle, color: '#10b981' },
          { label: 'En Operación', value: (byStatus.en_lavado || 0) + (byStatus.en_transito || 0) + (byStatus.asignado || 0), icon: Truck, color: '#3b82f6' },
          { label: 'Requieren Atención', value: (byStatus.en_reparacion || 0) + (byStatus.no_conforme || 0) + (byStatus.perdido || 0), icon: Wrench, color: '#ef4444' },
        ].map(({ label, value, icon: Icon, color }) => (
          <div key={label} className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
            <div className="flex items-center justify-between mb-2">
              <Icon size={18} style={{ color }} />
              <span className="text-2xl font-bold font-mono" style={{ color }}>{value}</span>
            </div>
            <p className="text-[#6b7280] text-xs">{label}</p>
          </div>
        ))}
      </div>

      <div className="grid grid-cols-1 xl:grid-cols-2 gap-4">
        {/* Workshops */}
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl">
          <div className="px-4 py-3 border-b border-[#1f2937]">
            <h3 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2">
              <Wrench size={15} className="text-[#f97316]" />
              Talleres Aliados
            </h3>
          </div>
          <div className="p-4 space-y-4">
            {workshops.map(w => (
              <div key={w.id} className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg p-3 space-y-2">
                <div className="flex items-start justify-between gap-2">
                  <div>
                    <p className="text-[#f9fafb] text-sm font-medium">{w.name}</p>
                    <p className="text-[#6b7280] text-xs font-mono">{w.id}</p>
                  </div>
                  <div className="flex items-center gap-1.5 flex-shrink-0">
                    {w.certified
                      ? <span className="flex items-center gap-1 text-emerald-400 text-xs bg-emerald-900/30 border border-emerald-900 px-2 py-0.5 rounded">
                          <CheckCircle size={10} /> Cert.
                        </span>
                      : <span className="flex items-center gap-1 text-amber-400 text-xs bg-amber-900/30 border border-amber-900 px-2 py-0.5 rounded">
                          <XCircle size={10} /> Pend.
                        </span>
                    }
                  </div>
                </div>
                <CapacityBar current={w.current} capacity={w.capacity} />
                <div className="flex items-center justify-between text-xs">
                  <span className="text-[#6b7280]">Próximo slot</span>
                  <span className="text-[#9ca3af] font-mono">{w.nextSlot}</span>
                </div>
              </div>
            ))}
          </div>
        </div>

        {/* Inventory summary */}
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl">
          <div className="px-4 py-3 border-b border-[#1f2937]">
            <h3 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2">
              <Package size={15} className="text-[#f97316]" />
              Resumen de Inventario
            </h3>
          </div>
          <div className="overflow-x-auto">
            <table className="w-full">
              <thead>
                <tr className="border-b border-[#1f2937]">
                  {['Producto', 'SKU', 'Edomex', 'Qro', 'Puebla', 'Total', 'Estado'].map(h => (
                    <th key={h} className="px-3 py-2 text-left text-xs text-[#6b7280] font-semibold uppercase tracking-wide">{h}</th>
                  ))}
                </tr>
              </thead>
              <tbody className="divide-y divide-[#1f2937]">
                {inventorySummary.map(p => (
                  <tr key={p.id} className="hover:bg-white/[0.02]">
                    <td className="px-3 py-2.5">
                      <p className="text-[#f9fafb] text-xs font-medium leading-tight max-w-32 truncate">{p.name}</p>
                    </td>
                    <td className="px-3 py-2.5">
                      <span className="text-[#6b7280] text-xs font-mono">{p.sku}</span>
                    </td>
                    {['bodega_edomex', 'bodega_queretaro', 'bodega_puebla'].map(wh => (
                      <td key={wh} className="px-3 py-2.5 text-center">
                        <span className={`text-xs font-mono ${
                          p.stock[wh] === 999 ? 'text-[#6b7280]' :
                          p.stock[wh] < 10 ? 'text-red-400' :
                          p.stock[wh] < 20 ? 'text-amber-400' : 'text-emerald-400'
                        }`}>
                          {p.stock[wh] === 999 ? '∞' : p.stock[wh]}
                        </span>
                      </td>
                    ))}
                    <td className="px-3 py-2.5 text-center">
                      <span className="text-[#f9fafb] text-xs font-mono font-bold">
                        {p.category === 'Servicios' ? '∞' : p.totalStock}
                      </span>
                    </td>
                    <td className="px-3 py-2.5">
                      {p.isLow
                        ? <span className="text-red-400 text-xs flex items-center gap-1"><XCircle size={10} />Bajo</span>
                        : <span className="text-emerald-400 text-xs flex items-center gap-1"><CheckCircle size={10} />OK</span>
                      }
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>

      {/* QR Feed */}
      <QRFeed activity={activity} />
    </div>
  )
}
