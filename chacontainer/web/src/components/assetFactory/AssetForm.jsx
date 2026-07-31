import { Plus, Trash2, Upload, X } from 'lucide-react'

const inputClass =
  'w-full bg-[#0a0f1a] border border-[#374151] rounded-lg px-3 py-2 text-[#f9fafb] text-sm focus:outline-none focus:border-[#f97316]'

function Field({ label, children }) {
  return (
    <div>
      <label className="text-xs text-[#9ca3af] block mb-1.5">{label}</label>
      {children}
    </div>
  )
}

export default function AssetForm({ assetType, value, onChange }) {
  const fields = assetType.fields || {}

  function set(key, val) {
    onChange({ ...value, [key]: val })
  }

  function setDatoClave(index, key, val) {
    const next = value.datosClave.slice()
    next[index] = { ...next[index], [key]: val }
    onChange({ ...value, datosClave: next })
  }

  function addDatoClave() {
    onChange({ ...value, datosClave: [...value.datosClave, { label: '', value: '' }] })
  }

  function removeDatoClave(index) {
    onChange({ ...value, datosClave: value.datosClave.filter((_, i) => i !== index) })
  }

  function handleImage(e) {
    const file = e.target.files?.[0]
    if (!file) return
    const reader = new FileReader()
    reader.onload = () => set('imagen', reader.result)
    reader.readAsDataURL(file)
  }

  return (
    <div className="space-y-4">
      <Field label="Título">
        <input
          className={inputClass}
          value={value.titulo}
          onChange={e => set('titulo', e.target.value)}
          placeholder="Título de la pieza"
        />
      </Field>

      {fields.subtitulo && (
        <Field label="Texto / subtítulo">
          <textarea
            className={`${inputClass} resize-none`}
            rows={3}
            value={value.subtitulo}
            onChange={e => set('subtitulo', e.target.value)}
            placeholder="Texto de apoyo, descripción o mensaje principal"
          />
        </Field>
      )}

      {fields.autor && (
        <Field label="Autor / atribución">
          <input
            className={inputClass}
            value={value.autor}
            onChange={e => set('autor', e.target.value)}
            placeholder="— Nombre, cargo"
          />
        </Field>
      )}

      {fields.cta && (
        <Field label="Texto del botón (CTA)">
          <input
            className={inputClass}
            value={value.cta}
            onChange={e => set('cta', e.target.value)}
            placeholder="Ej: Solicitar cotización"
          />
        </Field>
      )}

      {fields.datosClave && (
        <div>
          <div className="flex items-center justify-between mb-1.5">
            <label className="text-xs text-[#9ca3af]">{fields.datosClaveLabel || 'Datos clave'}</label>
            <button
              type="button"
              onClick={addDatoClave}
              className="flex items-center gap-1 text-xs text-[#f97316] hover:text-[#ea6c0a]"
            >
              <Plus size={13} /> Agregar
            </button>
          </div>
          <div className="space-y-2">
            {value.datosClave.map((d, i) => (
              <div key={i} className="flex gap-2 items-center">
                <input
                  className={`${inputClass} flex-1`}
                  value={d.label}
                  onChange={e => setDatoClave(i, 'label', e.target.value)}
                  placeholder="Título"
                />
                <input
                  className={`${inputClass} flex-1`}
                  value={d.value}
                  onChange={e => setDatoClave(i, 'value', e.target.value)}
                  placeholder="Detalle / valor"
                />
                <button
                  type="button"
                  onClick={() => removeDatoClave(i)}
                  className="text-[#6b7280] hover:text-[#ef4444] p-1.5 flex-shrink-0"
                >
                  <Trash2 size={14} />
                </button>
              </div>
            ))}
            {value.datosClave.length === 0 && (
              <p className="text-[#6b7280] text-xs italic">Sin filas todavía. Agrega al menos una.</p>
            )}
          </div>
        </div>
      )}

      {fields.imagen && (
        <Field label="Imagen (opcional)">
          {value.imagen ? (
            <div className="relative">
              <img src={value.imagen} alt="" className="w-full h-28 object-cover rounded-lg border border-[#1f2937]" />
              <button
                type="button"
                onClick={() => set('imagen', null)}
                className="absolute top-1.5 right-1.5 bg-black/60 hover:bg-black/80 rounded-full p-1"
              >
                <X size={13} className="text-white" />
              </button>
            </div>
          ) : (
            <label className="flex items-center justify-center gap-2 border border-dashed border-[#374151] rounded-lg py-4 text-xs text-[#6b7280] hover:border-[#f97316] hover:text-[#f97316] cursor-pointer transition-colors">
              <Upload size={14} />
              Subir imagen
              <input type="file" accept="image/*" className="hidden" onChange={handleImage} />
            </label>
          )}
        </Field>
      )}
    </div>
  )
}
