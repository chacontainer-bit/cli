import { useEffect, useState, useCallback } from 'react'
import { Warehouse, Plus, Loader2, X, MapPin } from 'lucide-react'
import { api, ApiError } from '../lib/api'

function NewPlantForm({ onCreated, onClose }) {
  const [form, setForm] = useState({ code: '', name: '', city: '', state: '' })
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState(null)

  async function handleSubmit(e) {
    e.preventDefault()
    setSaving(true)
    setError(null)
    try {
      await api.createPlant({
        code: form.code,
        name: form.name,
        address: { city: form.city, state: form.state, country: 'MX' },
        capacity: { total_slots: 100, used_slots: 0, max_weight_kg: 20000 },
      })
      onCreated()
      onClose()
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'No se pudo crear la planta')
    } finally {
      setSaving(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="bg-[#111827] border border-[#1f2937] rounded-xl p-4 space-y-3">
      <div className="flex items-center justify-between">
        <h3 className="text-sm font-semibold text-[#f9fafb]">Nueva planta</h3>
        <button type="button" onClick={onClose} className="text-[#6b7280] hover:text-[#f9fafb]"><X size={16} /></button>
      </div>
      <div className="grid grid-cols-2 gap-3">
        <input required placeholder="Código (ej. PLT-MTY)" value={form.code} onChange={e => setForm(f => ({ ...f, code: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]" />
        <input required placeholder="Nombre" value={form.name} onChange={e => setForm(f => ({ ...f, name: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]" />
        <input placeholder="Ciudad" value={form.city} onChange={e => setForm(f => ({ ...f, city: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]" />
        <input placeholder="Estado" value={form.state} onChange={e => setForm(f => ({ ...f, state: e.target.value }))}
          className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]" />
      </div>
      {error && <p className="text-red-400 text-xs">{error}</p>}
      <button type="submit" disabled={saving} className="bg-[#f97316] hover:bg-[#ea6a0d] disabled:opacity-60 text-white text-sm font-semibold px-4 py-2 rounded-lg flex items-center gap-2">
        {saving && <Loader2 size={14} className="animate-spin" />} Crear planta
      </button>
    </form>
  )
}

export default function Plants() {
  const [plants, setPlants] = useState([])
  const [loading, setLoading] = useState(true)
  const [showForm, setShowForm] = useState(false)

  const load = useCallback(async () => {
    setLoading(true)
    try {
      const res = await api.listPlants()
      setPlants(res.data || [])
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    let cancelled = false
    api.listPlants().then((res) => {
      if (cancelled) return
      setPlants(res.data || [])
      setLoading(false)
    })
    return () => { cancelled = true }
  }, [])

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-xl font-bold text-[#f9fafb] flex items-center gap-2">
            <Warehouse size={20} className="text-[#f97316]" /> Plantas
          </h1>
          <p className="text-[#6b7280] text-sm">Facilidades físicas y su capacidad</p>
        </div>
        <button onClick={() => setShowForm(s => !s)}
          className="flex items-center gap-1.5 text-xs font-medium px-3 py-2 rounded-lg bg-[#f97316] text-white hover:bg-[#ea6a0d]">
          <Plus size={14} /> Nueva planta
        </button>
      </div>

      {showForm && <NewPlantForm onCreated={load} onClose={() => setShowForm(false)} />}

      {loading ? (
        <div className="py-12 text-center text-[#6b7280] text-sm flex items-center justify-center gap-2"><Loader2 size={16} className="animate-spin" /> Cargando…</div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {plants.map(p => {
            const pct = p.capacity?.total_slots ? Math.round((p.capacity.used_slots / p.capacity.total_slots) * 100) : 0
            return (
              <div key={p.id} className="bg-[#111827] border border-[#1f2937] rounded-xl p-4 space-y-3">
                <div className="flex items-start justify-between">
                  <div>
                    <p className="text-[#f9fafb] text-sm font-semibold">{p.name}</p>
                    <p className="text-[#6b7280] text-xs font-mono">{p.code}</p>
                  </div>
                  <span className={`text-xs px-2 py-0.5 rounded-full border ${p.status === 'active' ? 'bg-emerald-900/30 text-emerald-400 border-emerald-900' : 'bg-gray-900/40 text-gray-400 border-gray-700'}`}>
                    {p.status}
                  </span>
                </div>
                {p.address?.city && (
                  <p className="text-[#9ca3af] text-xs flex items-center gap-1">
                    <MapPin size={11} /> {p.address.city}{p.address.state ? `, ${p.address.state}` : ''}
                  </p>
                )}
                <div>
                  <div className="flex justify-between text-xs mb-1">
                    <span className="text-[#9ca3af] font-mono">{p.capacity?.used_slots || 0}/{p.capacity?.total_slots || 0}</span>
                    <span className="text-[#6b7280]">{pct}%</span>
                  </div>
                  <div className="h-2 bg-[#1f2937] rounded-full overflow-hidden">
                    <div className="h-full bg-[#f97316] rounded-full transition-all" style={{ width: `${pct}%` }} />
                  </div>
                </div>
              </div>
            )
          })}
          {plants.length === 0 && (
            <p className="col-span-full text-center text-[#6b7280] text-sm py-8">Sin plantas registradas todavía</p>
          )}
        </div>
      )}
    </div>
  )
}
