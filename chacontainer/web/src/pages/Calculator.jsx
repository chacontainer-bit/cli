import { useState, useMemo } from 'react'
import {
  BarChart, Bar, XAxis, YAxis, Tooltip, ResponsiveContainer,
  PieChart, Pie, Cell, Legend
} from 'recharts'
import { Calculator, Leaf, DollarSign, TrendingUp, Clock } from 'lucide-react'

const containerTypes = [
  { id: 'ibc1000', label: 'IBC 1000L', priceNew: 850000, priceReuse: 285000, co2New: 42, co2Reuse: 12.4, lifespan: 10 },
  { id: 'ibc600', label: 'IBC 600L', priceNew: 560000, priceReuse: 195000, co2New: 28, co2Reuse: 8.2, lifespan: 10 },
  { id: 'tam200', label: 'Tambor 200L HDPE', priceNew: 120000, priceReuse: 45000, co2New: 8.5, co2Reuse: 2.1, lifespan: 8 },
  { id: 'tam220', label: 'Tambor 220L Metálico', priceNew: 145000, priceReuse: 52000, co2New: 12.0, co2Reuse: 1.8, lifespan: 12 },
]

const CustomTooltip = ({ active, payload, label }) => {
  if (active && payload && payload.length) {
    return (
      <div className="bg-[#1a2235] border border-[#1f2937] rounded-lg p-3">
        <p className="text-[#9ca3af] text-xs mb-1">{label}</p>
        {payload.map((p, i) => (
          <p key={i} className="text-sm font-mono" style={{ color: p.fill || p.color }}>
            {p.name}: {typeof p.value === 'number' ? p.value.toLocaleString('es-MX') : p.value}
          </p>
        ))}
      </div>
    )
  }
  return null
}

function ResultCard({ icon: Icon, label, value, unit, color, sub }) {
  return (
    <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
      <div className="flex items-center gap-2 mb-3">
        <div className="w-8 h-8 rounded-lg flex items-center justify-center" style={{ backgroundColor: `${color}20` }}>
          <Icon size={16} style={{ color }} />
        </div>
        <span className="text-xs text-[#6b7280]">{label}</span>
      </div>
      <p className="text-2xl font-bold font-mono" style={{ color }}>{value}</p>
      <p className="text-[#6b7280] text-xs mt-0.5">{unit}</p>
      {sub && <p className="text-[#9ca3af] text-xs mt-2 border-t border-[#1f2937] pt-2">{sub}</p>}
    </div>
  )
}

export default function CalculatorPage() {
  const [form, setForm] = useState({
    containerType: 'ibc1000',
    quantity: 50,
    annualTrips: 6,
    yearsAnalysis: 3,
  })

  function set(key, val) {
    setForm(prev => ({ ...prev, [key]: val }))
  }

  const result = useMemo(() => {
    const type = containerTypes.find(t => t.id === form.containerType)
    if (!type) return null

    const qty = Number(form.quantity) || 0
    const trips = Number(form.annualTrips) || 0
    const years = Number(form.yearsAnalysis) || 1

    // Cost new containers (replaced every trip cycle, paid per use)
    const annualCostNew = qty * trips * type.priceNew
    const totalCostNew = annualCostNew * years

    // Cost reutilizable (purchase amortized + service cost per wash)
    const purchaseCost = qty * type.priceReuse
    const washCostPerTrip = qty * trips * 28000 // avg wash cost
    const annualCostReuse = washCostPerTrip + (purchaseCost / type.lifespan)
    const totalCostReuse = annualCostReuse * years + purchaseCost

    const totalSavings = totalCostNew - totalCostReuse
    const annualSavings = annualCostNew - annualCostReuse
    const roiPct = totalCostReuse > 0 ? ((totalSavings / totalCostReuse) * 100) : 0
    const paybackMonths = annualSavings > 0 ? (purchaseCost / annualSavings) * 12 : 0

    // CO2
    const totalUsesNew = qty * trips * years
    const co2New = totalUsesNew * type.co2New
    const co2Reuse = qty * type.co2Reuse + (qty * trips * years * 1.2) // transport + wash
    const co2Saved = co2New - co2Reuse

    return {
      type,
      totalCostNew,
      totalCostReuse,
      totalSavings,
      annualSavings,
      roiPct,
      paybackMonths,
      co2New,
      co2Reuse,
      co2Saved,
    }
  }, [form])

  const barData = result ? [
    { name: 'Nuevo', costo: Math.round(result.totalCostNew / 1000), fill: '#ef4444' },
    { name: 'Reutilizable', costo: Math.round(result.totalCostReuse / 1000), fill: '#10b981' },
  ] : []

  const co2Data = result ? [
    { name: 'CO₂ Emitido (reuse)', value: Math.round(result.co2Reuse), fill: '#3b82f6' },
    { name: 'CO₂ Evitado', value: Math.round(result.co2Saved), fill: '#10b981' },
  ] : []

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-xl font-bold text-[#f9fafb] flex items-center gap-2">
          <Calculator size={20} className="text-[#f97316]" />
          Calculadora ROI & CO₂
        </h1>
        <p className="text-[#6b7280] text-sm mt-1">Calcule el ahorro económico y ambiental de cambiar a contenedores reutilizables</p>
      </div>

      <div className="grid grid-cols-1 xl:grid-cols-2 gap-6">
        {/* Form */}
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-5 space-y-5">
          <h2 className="text-sm font-semibold text-[#f9fafb] border-b border-[#1f2937] pb-3">
            Parámetros de Cálculo
          </h2>

          <div>
            <label className="text-xs text-[#9ca3af] block mb-2">Tipo de Contenedor</label>
            <div className="grid grid-cols-2 gap-2">
              {containerTypes.map(t => (
                <button key={t.id} onClick={() => set('containerType', t.id)}
                  className={`p-2.5 rounded-lg text-xs border text-left transition-all ${
                    form.containerType === t.id
                      ? 'bg-[#f97316]/10 border-[#f97316]/50 text-[#f97316]'
                      : 'bg-[#0a0f1a] border-[#1f2937] text-[#9ca3af] hover:border-[#374151]'
                  }`}>
                  <div className="font-semibold">{t.label}</div>
                  <div className="opacity-70 mt-0.5">${t.priceReuse.toLocaleString('es-MX')}/un</div>
                </button>
              ))}
            </div>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="text-xs text-[#9ca3af] block mb-1.5">Cantidad de Contenedores</label>
              <input type="number" min="1" value={form.quantity} onChange={e => set('quantity', e.target.value)}
                className="w-full bg-[#0a0f1a] border border-[#374151] rounded-lg px-3 py-2 text-[#f9fafb] text-sm focus:outline-none focus:border-[#f97316] font-mono" />
            </div>
            <div>
              <label className="text-xs text-[#9ca3af] block mb-1.5">Viajes/Ciclos Anuales</label>
              <input type="number" min="1" value={form.annualTrips} onChange={e => set('annualTrips', e.target.value)}
                className="w-full bg-[#0a0f1a] border border-[#374151] rounded-lg px-3 py-2 text-[#f9fafb] text-sm focus:outline-none focus:border-[#f97316] font-mono" />
            </div>
          </div>

          <div>
            <label className="text-xs text-[#9ca3af] block mb-1.5">Horizonte de Análisis (años)</label>
            <div className="flex gap-2">
              {[1, 2, 3, 5, 10].map(y => (
                <button key={y} onClick={() => set('yearsAnalysis', y)}
                  className={`flex-1 py-2 rounded-lg text-xs font-mono font-semibold border transition-all ${
                    form.yearsAnalysis === y
                      ? 'bg-[#f97316] border-[#f97316] text-white'
                      : 'bg-[#0a0f1a] border-[#374151] text-[#9ca3af] hover:border-[#6b7280]'
                  }`}>
                  {y}a
                </button>
              ))}
            </div>
          </div>

          {result && (
            <div className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg p-3 space-y-2">
              <p className="text-xs text-[#6b7280] font-semibold uppercase tracking-wide">Tipo seleccionado</p>
              <div className="grid grid-cols-2 gap-2 text-xs">
                <div><span className="text-[#6b7280]">Precio nuevo: </span>
                  <span className="text-[#f9fafb] font-mono">${result.type.priceNew.toLocaleString('es-MX')}</span></div>
                <div><span className="text-[#6b7280]">Precio reuse: </span>
                  <span className="text-[#f9fafb] font-mono">${result.type.priceReuse.toLocaleString('es-MX')}</span></div>
                <div><span className="text-[#6b7280]">CO₂ nuevo: </span>
                  <span className="text-red-400 font-mono">{result.type.co2New} kg/un</span></div>
                <div><span className="text-[#6b7280]">CO₂ reuse: </span>
                  <span className="text-emerald-400 font-mono">{result.type.co2Reuse} kg/un</span></div>
              </div>
            </div>
          )}
        </div>

        {/* Results */}
        {result && (
          <div className="space-y-4">
            <h2 className="text-sm font-semibold text-[#f9fafb]">Resultados ({form.yearsAnalysis} año{form.yearsAnalysis > 1 ? 's' : ''})</h2>
            <div className="grid grid-cols-2 gap-3">
              <ResultCard
                icon={DollarSign}
                label="Ahorro Total"
                value={`$${(result.totalSavings / 1000000).toFixed(1)}M`}
                unit="MXN acumulado"
                color="#10b981"
                sub={`$${Math.round(result.annualSavings / 1000).toLocaleString('es-MX')}K / año`}
              />
              <ResultCard
                icon={TrendingUp}
                label="ROI"
                value={`${Math.round(result.roiPct)}%`}
                unit="retorno sobre inversión"
                color="#f97316"
                sub={`Payback: ${Math.round(result.paybackMonths)} meses`}
              />
              <ResultCard
                icon={Leaf}
                label="CO₂ Evitado"
                value={`${Math.round(result.co2Saved / 1000).toLocaleString('es-MX')}t`}
                unit="toneladas CO₂ eq."
                color="#10b981"
                sub={`≈ ${Math.round(result.co2Saved / 200)} árboles/año`}
              />
              <ResultCard
                icon={Clock}
                label="Payback"
                value={`${Math.round(result.paybackMonths)}`}
                unit="meses para recuperar inversión"
                color="#3b82f6"
                sub={`Inversión inicial: $${Math.round(Number(form.quantity) * result.type.priceReuse / 1000).toLocaleString('es-MX')}K`}
              />
            </div>
          </div>
        )}
      </div>

      {/* Charts */}
      {result && (
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
          <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
            <h3 className="text-sm font-semibold text-[#f9fafb] mb-4">Comparación de Costo Total (MXN miles)</h3>
            <ResponsiveContainer width="100%" height={200}>
              <BarChart data={barData} barCategoryGap="40%">
                <XAxis dataKey="name" tick={{ fill: '#9ca3af', fontSize: 12 }} axisLine={false} tickLine={false} />
                <YAxis tick={{ fill: '#6b7280', fontSize: 11 }} axisLine={false} tickLine={false} />
                <Tooltip content={<CustomTooltip />} cursor={{ fill: 'rgba(255,255,255,0.03)' }} />
                <Bar dataKey="costo" name="Costo (K MXN)" radius={[4, 4, 0, 0]}>
                  {barData.map((entry, index) => (
                    <Cell key={index} fill={entry.fill} />
                  ))}
                </Bar>
              </BarChart>
            </ResponsiveContainer>
            <div className="mt-3 p-3 bg-emerald-900/20 border border-emerald-900/40 rounded-lg">
              <p className="text-emerald-400 text-sm font-bold font-mono">
                Ahorro: ${(result.totalSavings / 1000).toLocaleString('es-MX')}K MXN
              </p>
              <p className="text-[#9ca3af] text-xs mt-0.5">
                Los contenedores reutilizables cuestan {Math.round((result.totalCostReuse / result.totalCostNew) * 100)}% del precio de nuevos
              </p>
            </div>
          </div>

          <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
            <h3 className="text-sm font-semibold text-[#f9fafb] mb-4">Huella de Carbono (ton CO₂ eq.)</h3>
            <ResponsiveContainer width="100%" height={200}>
              <PieChart>
                <Pie data={co2Data} cx="50%" cy="50%" innerRadius={55} outerRadius={80}
                  paddingAngle={3} dataKey="value">
                  {co2Data.map((entry, index) => (
                    <Cell key={index} fill={entry.fill} />
                  ))}
                </Pie>
                <Tooltip
                  formatter={(v) => [`${v} ton CO₂`, '']}
                  contentStyle={{ backgroundColor: '#1a2235', border: '1px solid #1f2937', borderRadius: '8px', color: '#f9fafb', fontSize: '12px' }}
                />
                <Legend formatter={(v) => <span style={{ color: '#9ca3af', fontSize: '11px' }}>{v}</span>} />
              </PieChart>
            </ResponsiveContainer>
            <div className="mt-2 p-3 bg-[#0a0f1a] rounded-lg">
              <div className="flex justify-between text-xs">
                <span className="text-[#6b7280]">Reducción CO₂</span>
                <span className="text-emerald-400 font-mono font-bold">
                  {result.co2New > 0 ? Math.round(((result.co2New - result.co2Reuse) / result.co2New) * 100) : 0}%
                </span>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}
