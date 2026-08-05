import { useEffect, useState, useCallback } from 'react'
import { Settings, Truck, Users, Plus, Loader2, X, ArrowRight } from 'lucide-react'
import { api, ApiError } from '../lib/api'

const nextStatus = {
  draft: 'confirmed', confirmed: 'in_transit', in_transit: 'delivered',
}
const statusColor = {
  draft: 'text-gray-400', confirmed: 'text-blue-400', in_transit: 'text-orange-400',
  at_customs: 'text-amber-400', delivered: 'text-emerald-400', cancelled: 'text-red-400', exception: 'text-red-400',
}

function NewShipmentForm({ plants, clients, onCreated, onClose }) {
  const [form, setForm] = useState({ client_id: '', origin_plant_id: '', dest_plant_id: '' })
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState(null)

  async function handleSubmit(e) {
    e.preventDefault()
    setSaving(true)
    setError(null)
    try {
      await api.createShipment(form)
      onCreated()
      onClose()
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'No se pudo crear el envío')
    } finally {
      setSaving(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="bg-[#111827] border border-[#1f2937] rounded-xl p-4 space-y-3">
      <div className="flex items-center justify-between">
        <h3 className="text-sm font-semibold text-[#f9fafb]">Nuevo envío</h3>
        <button type="button" onClick={onClose} className="text-[#6b7280] hover:text-[#f9fafb]"><X size={16} /></button>
      </div>
      <div className="grid grid-cols-3 gap-3">
        <select required value={form.client_id} onChange={e => setForm(f => ({ ...f, client_id: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] focus:outline-none focus:border-[#f97316]">
          <option value="">Cliente…</option>
          {clients.map(c => <option key={c.id} value={c.id}>{c.name}</option>)}
        </select>
        <select required value={form.origin_plant_id} onChange={e => setForm(f => ({ ...f, origin_plant_id: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] focus:outline-none focus:border-[#f97316]">
          <option value="">Planta origen…</option>
          {plants.map(p => <option key={p.id} value={p.id}>{p.name}</option>)}
        </select>
        <select value={form.dest_plant_id} onChange={e => setForm(f => ({ ...f, dest_plant_id: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] focus:outline-none focus:border-[#f97316]">
          <option value="">Planta destino (opcional)…</option>
          {plants.map(p => <option key={p.id} value={p.id}>{p.name}</option>)}
        </select>
      </div>
      {error && <p className="text-red-400 text-xs">{error}</p>}
      <button type="submit" disabled={saving} className="bg-[#f97316] hover:bg-[#ea6a0d] disabled:opacity-60 text-white text-sm font-semibold px-4 py-2 rounded-lg flex items-center gap-2">
        {saving && <Loader2 size={14} className="animate-spin" />} Crear envío
      </button>
    </form>
  )
}

function NewClientForm({ onCreated, onClose }) {
  const [form, setForm] = useState({ code: '', name: '', type: 'shipper' })
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState(null)

  async function handleSubmit(e) {
    e.preventDefault()
    setSaving(true)
    setError(null)
    try {
      await api.createClient(form)
      onCreated()
      onClose()
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'No se pudo crear el cliente')
    } finally {
      setSaving(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="bg-[#111827] border border-[#1f2937] rounded-xl p-4 space-y-3">
      <div className="flex items-center justify-between">
        <h3 className="text-sm font-semibold text-[#f9fafb]">Nuevo cliente</h3>
        <button type="button" onClick={onClose} className="text-[#6b7280] hover:text-[#f9fafb]"><X size={16} /></button>
      </div>
      <div className="grid grid-cols-3 gap-3">
        <input required placeholder="Código" value={form.code} onChange={e => setForm(f => ({ ...f, code: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]" />
        <input required placeholder="Nombre" value={form.name} onChange={e => setForm(f => ({ ...f, name: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]" />
        <select value={form.type} onChange={e => setForm(f => ({ ...f, type: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] focus:outline-none focus:border-[#f97316]">
          {['shipper', 'receiver', 'forwarder', 'manufacturer', 'distributor'].map(t => <option key={t} value={t}>{t}</option>)}
        </select>
      </div>
      {error && <p className="text-red-400 text-xs">{error}</p>}
      <button type="submit" disabled={saving} className="bg-[#f97316] hover:bg-[#ea6a0d] disabled:opacity-60 text-white text-sm font-semibold px-4 py-2 rounded-lg flex items-center gap-2">
        {saving && <Loader2 size={14} className="animate-spin" />} Crear cliente
      </button>
    </form>
  )
}

export default function Backoffice() {
  const [shipments, setShipments] = useState([])
  const [clients, setClients] = useState([])
  const [plants, setPlants] = useState([])
  const [loading, setLoading] = useState(true)
  const [showForm, setShowForm] = useState(null) // 'shipment' | 'client' | null
  const [transitioning, setTransitioning] = useState(null)

  const load = useCallback(async () => {
    setLoading(true)
    try {
      const [sh, cl, pl] = await Promise.all([api.listShipments(), api.listClients(), api.listPlants()])
      setShipments(sh.data || [])
      setClients(cl.data || [])
      setPlants(pl.data || [])
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    let cancelled = false
    Promise.all([api.listShipments(), api.listClients(), api.listPlants()]).then(([sh, cl, pl]) => {
      if (cancelled) return
      setShipments(sh.data || [])
      setClients(cl.data || [])
      setPlants(pl.data || [])
      setLoading(false)
    })
    return () => { cancelled = true }
  }, [])

  async function advance(shipment) {
    const to = nextStatus[shipment.status]
    if (!to) return
    setTransitioning(shipment.id)
    try {
      await api.transitionShipment(shipment.id, { status: to })
      await load()
    } finally {
      setTransitioning(null)
    }
  }

  const byStatus = shipments.reduce((acc, s) => { acc[s.status] = (acc[s.status] || 0) + 1; return acc }, {})
  const clientById = Object.fromEntries(clients.map(c => [c.id, c.name]))

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-xl font-bold text-[#f9fafb] flex items-center gap-2">
            <Settings size={20} className="text-[#f97316]" /> Backoffice
          </h1>
          <p className="text-[#6b7280] text-sm">Envíos y cartera de clientes</p>
        </div>
        <div className="flex gap-2">
          <button onClick={() => setShowForm(showForm === 'client' ? null : 'client')}
            className="flex items-center gap-1.5 text-xs font-medium px-3 py-2 rounded-lg border border-[#1f2937] text-[#9ca3af] hover:text-[#f9fafb]">
            <Plus size={14} /> Cliente
          </button>
          <button onClick={() => setShowForm(showForm === 'shipment' ? null : 'shipment')}
            className="flex items-center gap-1.5 text-xs font-medium px-3 py-2 rounded-lg bg-[#f97316] text-white hover:bg-[#ea6a0d]">
            <Plus size={14} /> Envío
          </button>
        </div>
      </div>

      {showForm === 'shipment' && <NewShipmentForm plants={plants} clients={clients} onCreated={load} onClose={() => setShowForm(null)} />}
      {showForm === 'client' && <NewClientForm onCreated={load} onClose={() => setShowForm(null)} />}

      <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
        {[
          { label: 'Envíos totales', value: shipments.length, icon: Truck, color: '#f97316' },
          { label: 'En tránsito', value: byStatus.in_transit || 0, icon: ArrowRight, color: '#3b82f6' },
          { label: 'Entregados', value: byStatus.delivered || 0, icon: Truck, color: '#10b981' },
          { label: 'Clientes activos', value: clients.filter(c => c.status === 'active').length, icon: Users, color: '#eab308' },
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

      <div className="bg-[#111827] border border-[#1f2937] rounded-xl">
        <div className="px-4 py-3 border-b border-[#1f2937]">
          <h3 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2"><Truck size={15} className="text-[#f97316]" /> Envíos</h3>
        </div>
        {loading ? (
          <div className="py-10 text-center text-[#6b7280] text-sm flex items-center justify-center gap-2"><Loader2 size={16} className="animate-spin" /> Cargando…</div>
        ) : (
          <div className="divide-y divide-[#1f2937]">
            {shipments.map(s => (
              <div key={s.id} className="px-4 py-3 flex items-center justify-between gap-3">
                <div className="min-w-0 flex-1">
                  <p className="text-[#f9fafb] text-sm font-medium truncate">{s.reference}</p>
                  <p className="text-[#6b7280] text-xs truncate">{clientById[s.client_id] || s.client_id}</p>
                </div>
                <span className={`text-xs font-mono font-semibold ${statusColor[s.status]}`}>{s.status}</span>
                {nextStatus[s.status] && (
                  <button onClick={() => advance(s)} disabled={transitioning === s.id}
                    className="flex items-center gap-1 text-xs px-2.5 py-1 rounded-lg border border-[#1f2937] text-[#9ca3af] hover:text-[#f9fafb] disabled:opacity-50">
                    {transitioning === s.id ? <Loader2 size={12} className="animate-spin" /> : <ArrowRight size={12} />}
                    {nextStatus[s.status]}
                  </button>
                )}
              </div>
            ))}
            {shipments.length === 0 && <p className="text-center text-[#6b7280] text-sm py-8">Sin envíos aún</p>}
          </div>
        )}
      </div>

      <div className="bg-[#111827] border border-[#1f2937] rounded-xl">
        <div className="px-4 py-3 border-b border-[#1f2937]">
          <h3 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2"><Users size={15} className="text-[#f97316]" /> Clientes</h3>
        </div>
        <div className="overflow-x-auto">
          <table className="w-full">
            <thead>
              <tr className="border-b border-[#1f2937]">
                {['Nombre', 'Código', 'Tipo', 'Estado'].map(h => (
                  <th key={h} className="px-3 py-2 text-left text-xs text-[#6b7280] font-semibold uppercase tracking-wide">{h}</th>
                ))}
              </tr>
            </thead>
            <tbody className="divide-y divide-[#1f2937]">
              {clients.map(c => (
                <tr key={c.id} className="hover:bg-white/[0.02]">
                  <td className="px-3 py-2.5 text-[#f9fafb] text-xs font-medium">{c.name}</td>
                  <td className="px-3 py-2.5 text-[#6b7280] text-xs font-mono">{c.code}</td>
                  <td className="px-3 py-2.5 text-[#9ca3af] text-xs">{c.type}</td>
                  <td className="px-3 py-2.5 text-emerald-400 text-xs">{c.status}</td>
                </tr>
              ))}
              {clients.length === 0 && (
                <tr><td colSpan={4} className="text-center text-[#6b7280] text-sm py-8">Sin clientes aún</td></tr>
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
}
