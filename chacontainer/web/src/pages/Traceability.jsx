import { useEffect, useState, useCallback } from 'react'
import { Search, ChevronDown, ChevronUp, Clock, Plus, Loader2, ScanLine, X } from 'lucide-react'
import { api, ApiError } from '../lib/api'

const statusFilters = [
  { id: '', label: 'Todos' },
  { id: 'available', label: 'Disponible' },
  { id: 'in_use', label: 'En uso' },
  { id: 'in_transit', label: 'En tránsito' },
  { id: 'maintenance', label: 'Mantenimiento' },
  { id: 'retired', label: 'Retirado' },
  { id: 'lost', label: 'Perdido' },
]

const statusStyle = {
  available: { bg: 'bg-emerald-900/30', text: 'text-emerald-400', border: 'border-emerald-900', dot: '#10b981' },
  in_use: { bg: 'bg-blue-900/30', text: 'text-blue-400', border: 'border-blue-900', dot: '#3b82f6' },
  in_transit: { bg: 'bg-orange-900/30', text: 'text-orange-400', border: 'border-orange-900', dot: '#f97316' },
  maintenance: { bg: 'bg-amber-900/30', text: 'text-amber-400', border: 'border-amber-900', dot: '#eab308' },
  retired: { bg: 'bg-gray-900/40', text: 'text-gray-400', border: 'border-gray-700', dot: '#6b7280' },
  lost: { bg: 'bg-red-900/30', text: 'text-red-400', border: 'border-red-900', dot: '#ef4444' },
}

const assetTypes = ['container', 'pallet', 'forklift', 'rack', 'ibc', 'drum', 'trailer']

function StatusBadge({ status }) {
  const cfg = statusStyle[status] || statusStyle.available
  const label = statusFilters.find(f => f.id === status)?.label || status
  return (
    <span className={`inline-flex items-center gap-1.5 text-xs px-2 py-0.5 rounded-full border ${cfg.bg} ${cfg.text} ${cfg.border}`}>
      <span className="w-1.5 h-1.5 rounded-full flex-shrink-0" style={{ backgroundColor: cfg.dot }} />
      {label}
    </span>
  )
}

function EventTimeline({ assetId }) {
  const [events, setEvents] = useState(null)
  useEffect(() => {
    api.getAssetEvents(assetId).then(res => setEvents(res.data || [])).catch(() => setEvents([]))
  }, [assetId])

  return (
    <div className="border-t border-[#1f2937] bg-[#0a0f1a]">
      <div className="px-4 py-3">
        <h4 className="text-xs font-semibold text-[#9ca3af] uppercase tracking-wide mb-3 flex items-center gap-1.5">
          <Clock size={12} /> Historial de eventos
        </h4>
        {events === null && <p className="text-[#6b7280] text-xs">Cargando…</p>}
        {events?.length === 0 && <p className="text-[#6b7280] text-xs">Sin eventos registrados</p>}
        <div className="relative pl-4">
          {events && events.length > 0 && <div className="absolute left-0 top-0 bottom-0 w-px bg-[#1f2937]" />}
          {events?.map((ev) => (
            <div key={ev.id} className="relative mb-3 last:mb-0">
              <div className="absolute -left-4 top-1.5 w-2 h-2 rounded-full bg-[#f97316] border-2 border-[#0a0f1a]" />
              <div className="flex-1 min-w-0">
                <div className="flex items-center gap-2 mb-0.5">
                  <span className="text-[#f9fafb] text-xs font-semibold">{ev.event_type}</span>
                  <span className="text-[#6b7280] text-xs font-mono">{new Date(ev.occurred_at).toLocaleString('es-MX')}</span>
                </div>
                {ev.notes && <p className="text-[#9ca3af] text-xs">{ev.notes}</p>}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

function NewAssetForm({ onCreated, onClose }) {
  const [form, setForm] = useState({ name: '', code: '', type: 'container', plant_id: '' })
  const [plants, setPlants] = useState([])
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState(null)

  useEffect(() => { api.listPlants().then(res => setPlants(res.data || [])).catch(() => {}) }, [])

  async function handleSubmit(e) {
    e.preventDefault()
    setSaving(true)
    setError(null)
    try {
      await api.createAsset(form)
      onCreated()
      onClose()
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'No se pudo crear el activo')
    } finally {
      setSaving(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="bg-[#111827] border border-[#1f2937] rounded-xl p-4 space-y-3">
      <div className="flex items-center justify-between">
        <h3 className="text-sm font-semibold text-[#f9fafb]">Registrar activo</h3>
        <button type="button" onClick={onClose} className="text-[#6b7280] hover:text-[#f9fafb]"><X size={16} /></button>
      </div>
      <div className="grid grid-cols-2 gap-3">
        <input required placeholder="Nombre" value={form.name} onChange={e => setForm(f => ({ ...f, name: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]" />
        <input required placeholder="Código (ej. AST-0010)" value={form.code} onChange={e => setForm(f => ({ ...f, code: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]" />
        <select value={form.type} onChange={e => setForm(f => ({ ...f, type: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] focus:outline-none focus:border-[#f97316]">
          {assetTypes.map(t => <option key={t} value={t}>{t}</option>)}
        </select>
        <select required value={form.plant_id} onChange={e => setForm(f => ({ ...f, plant_id: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] focus:outline-none focus:border-[#f97316]">
          <option value="">Selecciona planta…</option>
          {plants.map(p => <option key={p.id} value={p.id}>{p.name}</option>)}
        </select>
      </div>
      {error && <p className="text-red-400 text-xs">{error}</p>}
      <button type="submit" disabled={saving} className="bg-[#f97316] hover:bg-[#ea6a0d] disabled:opacity-60 text-white text-sm font-semibold px-4 py-2 rounded-lg flex items-center gap-2">
        {saving && <Loader2 size={14} className="animate-spin" />} Crear activo
      </button>
    </form>
  )
}

function ScanForm({ onScanned, onClose }) {
  const [form, setForm] = useState({ qr_code: '', action: 'check_in' })
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState(null)
  const [result, setResult] = useState(null)

  async function handleSubmit(e) {
    e.preventDefault()
    setSaving(true)
    setError(null)
    setResult(null)
    try {
      const res = await api.scanAsset(form)
      setResult(res.asset)
      onScanned()
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'No se pudo procesar el escaneo')
    } finally {
      setSaving(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="bg-[#111827] border border-[#1f2937] rounded-xl p-4 space-y-3">
      <div className="flex items-center justify-between">
        <h3 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2"><ScanLine size={15} className="text-[#f97316]" /> Escanear QR</h3>
        <button type="button" onClick={onClose} className="text-[#6b7280] hover:text-[#f9fafb]"><X size={16} /></button>
      </div>
      <p className="text-[#6b7280] text-xs">Pega el código QR de un activo (columna "qr_code" en la tabla) para simular el escaneo.</p>
      <div className="grid grid-cols-2 gap-3">
        <input required placeholder="Código QR" value={form.qr_code} onChange={e => setForm(f => ({ ...f, qr_code: e.target.value }))}
          className="col-span-2 bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]" />
        <select value={form.action} onChange={e => setForm(f => ({ ...f, action: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] focus:outline-none focus:border-[#f97316]">
          {['check_in', 'check_out', 'transfer', 'inspection', 'maintenance'].map(a => <option key={a} value={a}>{a}</option>)}
        </select>
      </div>
      {error && <p className="text-red-400 text-xs">{error}</p>}
      {result && <p className="text-emerald-400 text-xs">Escaneado: {result.name} ({result.code}) → {result.status}</p>}
      <button type="submit" disabled={saving} className="bg-[#f97316] hover:bg-[#ea6a0d] disabled:opacity-60 text-white text-sm font-semibold px-4 py-2 rounded-lg flex items-center gap-2">
        {saving && <Loader2 size={14} className="animate-spin" />} Registrar escaneo
      </button>
    </form>
  )
}

export default function Traceability() {
  const [activeStatus, setActiveStatus] = useState('')
  const [search, setSearch] = useState('')
  const [expandedId, setExpandedId] = useState(null)
  const [assets, setAssets] = useState([])
  const [total, setTotal] = useState(0)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)
  const [showForm, setShowForm] = useState(null) // 'create' | 'scan' | null

  const load = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const params = {}
      if (activeStatus) params.status = activeStatus
      if (search) params.q = search
      const res = await api.listAssets(params)
      setAssets(res.data || [])
      setTotal(res.total || 0)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }, [activeStatus, search])

  useEffect(() => {
    let cancelled = false
    const params = {}
    if (activeStatus) params.status = activeStatus
    if (search) params.q = search
    api.listAssets(params).then(
      (res) => {
        if (cancelled) return
        setAssets(res.data || [])
        setTotal(res.total || 0)
        setError(null)
        setLoading(false)
      },
      (err) => {
        if (cancelled) return
        setError(err.message)
        setLoading(false)
      },
    )
    return () => { cancelled = true }
  }, [activeStatus, search])

  function toggleExpand(id) {
    setExpandedId(prev => prev === id ? null : id)
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-xl font-bold text-[#f9fafb]">Activos</h1>
          <p className="text-[#6b7280] text-sm">Trazabilidad de contenedores, pallets y equipo retornable</p>
        </div>
        <div className="flex gap-2">
          <button onClick={() => setShowForm(showForm === 'scan' ? null : 'scan')}
            className="flex items-center gap-1.5 text-xs font-medium px-3 py-2 rounded-lg border border-[#1f2937] text-[#9ca3af] hover:text-[#f9fafb]">
            <ScanLine size={14} /> Escanear QR
          </button>
          <button onClick={() => setShowForm(showForm === 'create' ? null : 'create')}
            className="flex items-center gap-1.5 text-xs font-medium px-3 py-2 rounded-lg bg-[#f97316] text-white hover:bg-[#ea6a0d]">
            <Plus size={14} /> Registrar activo
          </button>
        </div>
      </div>

      {showForm === 'create' && <NewAssetForm onCreated={load} onClose={() => setShowForm(null)} />}
      {showForm === 'scan' && <ScanForm onScanned={load} onClose={() => setShowForm(null)} />}

      <div className="flex flex-col sm:flex-row gap-3">
        <div className="relative flex-1">
          <Search size={14} className="absolute left-3 top-1/2 -translate-y-1/2 text-[#6b7280]" />
          <input
            type="text"
            placeholder="Buscar por nombre o código…"
            value={search}
            onChange={e => setSearch(e.target.value)}
            className="w-full bg-[#111827] border border-[#1f2937] rounded-lg pl-9 pr-3 py-2 text-sm text-[#f9fafb] placeholder-[#6b7280] focus:outline-none focus:border-[#f97316]"
          />
        </div>
      </div>

      <div className="flex gap-1.5 overflow-x-auto pb-1">
        {statusFilters.map(f => (
          <button key={f.id} onClick={() => setActiveStatus(f.id)}
            className={`flex-shrink-0 px-3 py-1.5 rounded-lg text-xs font-medium transition-colors ${
              activeStatus === f.id ? 'bg-[#f97316] text-white' : 'bg-[#111827] border border-[#1f2937] text-[#9ca3af] hover:text-[#f9fafb]'
            }`}>
            {f.label}
          </button>
        ))}
      </div>

      {error && <div className="bg-red-500/10 border border-red-500/30 rounded-xl p-4 text-red-400 text-sm">{error}</div>}

      <div className="bg-[#111827] border border-[#1f2937] rounded-xl overflow-hidden">
        <div className="grid grid-cols-12 gap-2 px-4 py-2 border-b border-[#1f2937] text-xs text-[#6b7280] font-semibold uppercase tracking-wide">
          <div className="col-span-3">Código / Nombre</div>
          <div className="col-span-2">Tipo</div>
          <div className="col-span-2">Estado</div>
          <div className="col-span-3">QR</div>
          <div className="col-span-2">Última actualización</div>
        </div>

        {loading && (
          <div className="py-12 text-center text-[#6b7280] flex items-center justify-center gap-2 text-sm">
            <Loader2 size={16} className="animate-spin" /> Cargando…
          </div>
        )}

        {!loading && (
          <div className="divide-y divide-[#1f2937]">
            {assets.map(a => {
              const isExpanded = expandedId === a.id
              return (
                <div key={a.id}>
                  <div onClick={() => toggleExpand(a.id)}
                    className="grid grid-cols-12 gap-2 px-4 py-3 cursor-pointer hover:bg-white/[0.02] transition-colors items-center">
                    <div className="col-span-3 min-w-0">
                      <p className="text-[#f97316] text-xs font-mono font-bold">{a.code}</p>
                      <p className="text-[#d1d5db] text-xs truncate">{a.name}</p>
                    </div>
                    <div className="col-span-2"><span className="text-[#d1d5db] text-xs">{a.type}</span></div>
                    <div className="col-span-2"><StatusBadge status={a.status} /></div>
                    <div className="col-span-3"><span className="text-[#6b7280] text-xs font-mono truncate block">{a.qr_code}</span></div>
                    <div className="col-span-2 flex items-center justify-between">
                      <span className="text-[#6b7280] text-xs font-mono">{new Date(a.updated_at).toLocaleDateString('es-MX')}</span>
                      {isExpanded ? <ChevronUp size={12} className="text-[#6b7280]" /> : <ChevronDown size={12} className="text-[#6b7280]" />}
                    </div>
                  </div>
                  {isExpanded && <EventTimeline assetId={a.id} />}
                </div>
              )
            })}
            {assets.length === 0 && (
              <div className="py-12 text-center text-[#6b7280]">
                <Search size={32} className="mx-auto mb-3 opacity-30" />
                <p className="text-sm">No se encontraron activos</p>
              </div>
            )}
          </div>
        )}
      </div>
      <p className="text-xs text-[#6b7280] font-mono">{assets.length} de {total} activos</p>
    </div>
  )
}
