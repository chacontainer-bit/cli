import { useState } from 'react'
import { Search, ChevronDown, ChevronUp, Clock, MapPin, User } from 'lucide-react'
import { containers, containerEvents, statusConfig } from '../data/mockData'

const statusFilters = [
  { id: 'todos', label: 'Todos' },
  { id: 'disponible', label: 'Disponible' },
  { id: 'en_lavado', label: 'En Lavado' },
  { id: 'en_transito', label: 'En Tránsito' },
  { id: 'asignado', label: 'Asignado' },
  { id: 'en_reparacion', label: 'En Reparación' },
  { id: 'no_conforme', label: 'No Conforme' },
  { id: 'perdido', label: 'Perdido' },
]

function StatusBadge({ status }) {
  const cfg = statusConfig[status] || { label: status, bg: 'bg-gray-900/40', text: 'text-gray-400', border: 'border-gray-700' }
  return (
    <span className={`inline-flex items-center gap-1.5 text-xs px-2 py-0.5 rounded-full border ${cfg.bg} ${cfg.text} ${cfg.border}`}>
      <span className="w-1.5 h-1.5 rounded-full flex-shrink-0" style={{ backgroundColor: cfg.color || '#6b7280' }} />
      {cfg.label}
    </span>
  )
}

function EventTimeline({ containerId }) {
  const events = containerEvents[containerId] || [
    { time: '2026-05-15 08:00', event: 'Sin eventos registrados', detail: 'No hay historial disponible', user: '-' }
  ]

  return (
    <div className="border-t border-[#1f2937] bg-[#0a0f1a]">
      <div className="px-4 py-3">
        <h4 className="text-xs font-semibold text-[#9ca3af] uppercase tracking-wide mb-3 flex items-center gap-1.5">
          <Clock size={12} />
          Historial de Eventos
        </h4>
        <div className="relative pl-4">
          <div className="absolute left-0 top-0 bottom-0 w-px bg-[#1f2937]" />
          {events.map((ev, i) => (
            <div key={i} className="relative mb-3 last:mb-0">
              <div className="absolute -left-4 top-1.5 w-2 h-2 rounded-full bg-[#f97316] border-2 border-[#0a0f1a]" />
              <div className="flex items-start gap-3">
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 mb-0.5">
                    <span className="text-[#f9fafb] text-xs font-semibold">{ev.event}</span>
                    <span className="text-[#6b7280] text-xs font-mono">{ev.time}</span>
                  </div>
                  <p className="text-[#9ca3af] text-xs">{ev.detail}</p>
                  <div className="flex items-center gap-1 mt-0.5">
                    <User size={10} className="text-[#6b7280]" />
                    <span className="text-[#6b7280] text-xs">{ev.user}</span>
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

// Distribución de activos por hub regional (Bajío / Centro de México).
function HubDistribution({ containers }) {
  const statusCounts = {}
  containers.forEach(c => {
    statusCounts[c.status] = (statusCounts[c.status] || 0) + 1
  })

  const hubs = [
    { label: 'Querétaro', count: containers.filter(c => c.location.includes('Querétaro')).length },
    { label: 'Puebla', count: containers.filter(c => c.location.includes('Puebla')).length },
    { label: 'Estado de México', count: containers.filter(c => c.location.includes('Estado de México')).length },
    { label: 'Villa de Reyes', count: containers.filter(c => c.location.includes('Villa de Reyes')).length },
    { label: 'Toluca', count: containers.filter(c => c.location.includes('Toluca')).length },
    { label: 'Guanajuato', count: containers.filter(c => c.location.includes('Guanajuato')).length },
    { label: 'San Luis Potosí', count: containers.filter(c => c.location.includes('San Luis Potosí')).length },
  ]
  const maxCount = Math.max(1, ...hubs.map(h => h.count))

  return (
    <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
      <h3 className="text-sm font-semibold text-[#f9fafb] mb-4 flex items-center gap-2">
        <MapPin size={14} className="text-[#f97316]" />
        Distribución por Hub
      </h3>

      <div className="space-y-2.5">
        {hubs.map(h => (
          <div key={h.label}>
            <div className="flex items-center justify-between text-xs mb-1">
              <span className="text-[#9ca3af]">{h.label}</span>
              <span className="text-[#f9fafb] font-mono">{h.count}</span>
            </div>
            <div className="h-1.5 bg-[#0a0f1a] rounded-full overflow-hidden">
              <div
                className="h-full bg-[#f97316] rounded-full"
                style={{ width: `${h.count > 0 ? Math.max(6, (h.count / maxCount) * 100) : 0}%` }}
              />
            </div>
          </div>
        ))}
      </div>

      <div className="mt-4 pt-3 border-t border-[#1f2937] space-y-1.5">
        {Object.entries(statusConfig).slice(0, 4).map(([key, cfg]) => (
          <div key={key} className="flex items-center justify-between text-xs">
            <div className="flex items-center gap-1.5">
              <div className="w-2 h-2 rounded-full" style={{ backgroundColor: cfg.color }} />
              <span className="text-[#9ca3af]">{cfg.label}</span>
            </div>
            <span className="text-[#f9fafb] font-mono">{statusCounts[key] || 0}</span>
          </div>
        ))}
      </div>
    </div>
  )
}

export default function Traceability() {
  const [activeStatus, setActiveStatus] = useState('todos')
  const [search, setSearch] = useState('')
  const [expandedId, setExpandedId] = useState(null)

  const filtered = containers.filter(c => {
    const matchStatus = activeStatus === 'todos' || c.status === activeStatus
    const matchSearch = !search || c.id.toLowerCase().includes(search.toLowerCase()) ||
      c.location.toLowerCase().includes(search.toLowerCase()) ||
      (c.client && c.client.toLowerCase().includes(search.toLowerCase()))
    return matchStatus && matchSearch
  })

  function toggleExpand(id) {
    setExpandedId(prev => prev === id ? null : id)
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-xl font-bold text-[#f9fafb]">Trazabilidad</h1>
          <p className="text-[#6b7280] text-sm">Seguimiento y estado de contenedores en tiempo real</p>
        </div>
        <div className="flex items-center gap-2 text-xs">
          <div className="w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse" />
          <span className="text-emerald-400 font-mono">LIVE</span>
        </div>
      </div>

      <div className="grid grid-cols-1 xl:grid-cols-4 gap-4">
        {/* Main content */}
        <div className="xl:col-span-3 space-y-4">
          {/* Filters */}
          <div className="flex flex-col sm:flex-row gap-3">
            <div className="relative flex-1">
              <Search size={14} className="absolute left-3 top-1/2 -translate-y-1/2 text-[#6b7280]" />
              <input
                type="text"
                placeholder="Buscar por ID, ubicación, cliente..."
                value={search}
                onChange={e => setSearch(e.target.value)}
                className="w-full bg-[#111827] border border-[#1f2937] rounded-lg pl-9 pr-3 py-2 text-sm text-[#f9fafb] placeholder-[#6b7280] focus:outline-none focus:border-[#f97316]"
              />
            </div>
          </div>

          {/* Status tabs */}
          <div className="flex gap-1.5 overflow-x-auto pb-1">
            {statusFilters.map(f => (
              <button key={f.id} onClick={() => setActiveStatus(f.id)}
                className={`flex-shrink-0 px-3 py-1.5 rounded-lg text-xs font-medium transition-colors ${
                  activeStatus === f.id
                    ? 'bg-[#f97316] text-white'
                    : 'bg-[#111827] border border-[#1f2937] text-[#9ca3af] hover:text-[#f9fafb]'
                }`}>
                {f.label}
                <span className="ml-1.5 font-mono opacity-70">
                  {f.id === 'todos' ? containers.length : containers.filter(c => c.status === f.id).length}
                </span>
              </button>
            ))}
          </div>

          {/* Container table */}
          <div className="bg-[#111827] border border-[#1f2937] rounded-xl overflow-hidden">
            <div className="grid grid-cols-12 gap-2 px-4 py-2 border-b border-[#1f2937] text-xs text-[#6b7280] font-semibold uppercase tracking-wide">
              <div className="col-span-2">ID</div>
              <div className="col-span-2">Tipo</div>
              <div className="col-span-2">Estado</div>
              <div className="col-span-3">Ubicación</div>
              <div className="col-span-1 text-center">Lavados</div>
              <div className="col-span-2">Último Evento</div>
            </div>

            <div className="divide-y divide-[#1f2937]">
              {filtered.map(container => {
                const isExpanded = expandedId === container.id
                return (
                  <div key={container.id}>
                    <div
                      onClick={() => toggleExpand(container.id)}
                      className="grid grid-cols-12 gap-2 px-4 py-3 cursor-pointer hover:bg-white/[0.02] transition-colors items-center"
                    >
                      <div className="col-span-2">
                        <span className="text-[#f97316] text-xs font-mono font-bold">{container.id}</span>
                      </div>
                      <div className="col-span-2">
                        <span className="text-[#d1d5db] text-xs">{container.type}</span>
                      </div>
                      <div className="col-span-2">
                        <StatusBadge status={container.status} />
                      </div>
                      <div className="col-span-3 flex items-center gap-1">
                        <MapPin size={10} className="text-[#6b7280] flex-shrink-0" />
                        <span className="text-[#9ca3af] text-xs truncate">{container.location}</span>
                      </div>
                      <div className="col-span-1 text-center">
                        <span className="text-[#9ca3af] text-xs font-mono">{container.washes}</span>
                      </div>
                      <div className="col-span-2 flex items-center justify-between">
                        <span className="text-[#6b7280] text-xs font-mono truncate">{container.lastEvent.split(' ')[1] || container.lastEvent}</span>
                        {isExpanded ? <ChevronUp size={12} className="text-[#6b7280]" /> : <ChevronDown size={12} className="text-[#6b7280]" />}
                      </div>
                    </div>
                    {isExpanded && <EventTimeline containerId={container.id} />}
                  </div>
                )
              })}
            </div>

            {filtered.length === 0 && (
              <div className="py-12 text-center text-[#6b7280]">
                <Search size={32} className="mx-auto mb-3 opacity-30" />
                <p className="text-sm">No se encontraron contenedores</p>
              </div>
            )}
          </div>

          <p className="text-xs text-[#6b7280] font-mono">{filtered.length} de {containers.length} contenedores</p>
        </div>

        {/* Hub distribution sidebar */}
        <div className="xl:col-span-1">
          <HubDistribution containers={containers} />
        </div>
      </div>
    </div>
  )
}
