import { useState } from 'react'
import {
  ClipboardList, PackageSearch, CircleAlert, CircleCheck, CircleDashed,
  Gauge, Lock, ShieldAlert
} from 'lucide-react'
import { ertPilot, ertAssets, ertMilestones, ertRoiGate } from '../data/mockData'

function AvailabilityBadge({ available }) {
  const cfg = {
    NO: { label: 'NO DISPONIBLE', bg: 'bg-red-900/40', text: 'text-red-400', border: 'border-red-700' },
    PARCIAL: { label: 'PARCIAL', bg: 'bg-amber-900/40', text: 'text-amber-400', border: 'border-amber-700' },
    SI: { label: 'DISPONIBLE', bg: 'bg-emerald-900/40', text: 'text-emerald-400', border: 'border-emerald-700' },
  }[available] || { label: available, bg: 'bg-gray-900/40', text: 'text-gray-400', border: 'border-gray-700' }

  return (
    <span className={`inline-flex items-center text-xs px-2 py-0.5 rounded-full border ${cfg.bg} ${cfg.text} ${cfg.border}`}>
      {cfg.label}
    </span>
  )
}

function GateRow({ label, target, status }) {
  return (
    <div className="flex items-start justify-between gap-3 py-2 border-b border-[#1f2937] last:border-b-0">
      <div className="min-w-0">
        <p className="text-[#f9fafb] text-xs font-semibold">{label}</p>
        <p className="text-[#6b7280] text-xs mt-0.5">{target}</p>
      </div>
      <span className="flex-shrink-0 inline-flex items-center gap-1 text-xs px-2 py-0.5 rounded-full border bg-gray-900/40 text-gray-400 border-gray-600">
        <Lock size={10} />
        {status}
      </span>
    </div>
  )
}

export default function CircuitoERT() {
  const [showAllAssets, setShowAllAssets] = useState(false)
  const completedCount = ertMilestones.filter(m => m.available === 'SI').length
  const partialCount = ertMilestones.filter(m => m.available === 'PARCIAL').length
  const visibleAssets = showAllAssets ? ertAssets : ertAssets.slice(0, 8)

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-start justify-between gap-4 flex-wrap">
        <div>
          <h1 className="text-xl font-bold text-[#f9fafb]">Circuito ERT — {ertPilot.client}</h1>
          <p className="text-[#6b7280] text-sm mt-0.5">
            {ertPilot.program} · {ertPilot.family} {ertPilot.measure} · Caso {ertPilot.caseCode}
          </p>
        </div>
        <div className="flex items-center gap-2 px-3 py-1.5 bg-amber-500/10 border border-amber-500/30 rounded-full">
          <CircleDashed size={13} className="text-amber-400" />
          <span className="text-amber-400 text-xs font-medium">Circuito {ertPilot.circuitStatus}</span>
        </div>
      </div>

      <p className="text-xs text-[#6b7280] bg-[#111827] border border-[#1f2937] rounded-lg px-3 py-2">
        {ertPilot.circuitStatusNote}
      </p>

      {/* KPI strip */}
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
          <p className="text-[#6b7280] text-xs">Universo reportado</p>
          <p className="text-2xl font-bold text-[#f9fafb] font-mono mt-1">{ertPilot.circuitPoolTotal.toLocaleString('es-MX')}</p>
          <p className="text-[#6b7280] text-xs mt-0.5">contenedores {ertPilot.measure}</p>
        </div>
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
          <p className="text-[#6b7280] text-xs">Muestra piloto</p>
          <p className="text-2xl font-bold text-[#f9fafb] font-mono mt-1">{ertPilot.sampleAssetsTarget}</p>
          <p className="text-[#6b7280] text-xs mt-0.5">activos con ID provisional</p>
        </div>
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
          <p className="text-[#6b7280] text-xs">Hitos del ciclo</p>
          <p className="text-2xl font-bold text-[#f9fafb] font-mono mt-1">{completedCount}/{ertMilestones.length}</p>
          <p className="text-[#6b7280] text-xs mt-0.5">{partialCount} parcial, resto pendiente</p>
        </div>
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
          <p className="text-[#6b7280] text-xs">Valor histórico</p>
          <p className="text-2xl font-bold text-[#f9fafb] font-mono mt-1">${ertPilot.historicalValueMXN.toLocaleString('es-MX')}</p>
          <p className="text-[#6b7280] text-xs mt-0.5">MXN · comercial, no costo</p>
        </div>
      </div>

      <div className="grid grid-cols-1 xl:grid-cols-3 gap-4">
        {/* Hitos obligatorios */}
        <div className="xl:col-span-2 bg-[#111827] border border-[#1f2937] rounded-xl overflow-hidden">
          <div className="px-4 py-3 border-b border-[#1f2937] flex items-center justify-between">
            <h2 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2">
              <ClipboardList size={15} className="text-[#f97316]" />
              Hitos obligatorios del primer ciclo
            </h2>
            <span className="text-xs font-mono text-[#6b7280]">09_CASO_ALPHA</span>
          </div>
          <div className="divide-y divide-[#1f2937] max-h-[480px] overflow-y-auto">
            {ertMilestones.map(m => (
              <div key={m.n} className="p-3 flex items-start gap-3">
                <span className="flex-shrink-0 w-6 h-6 rounded-full bg-[#0a0f1a] border border-[#1f2937] text-[#6b7280] text-xs font-mono flex items-center justify-center mt-0.5">
                  {m.n}
                </span>
                <div className="min-w-0 flex-1">
                  <div className="flex items-center justify-between gap-2">
                    <span className="text-[#f9fafb] text-xs font-semibold">{m.name}</span>
                    <AvailabilityBadge available={m.available} />
                  </div>
                  <p className="text-[#9ca3af] text-xs mt-1">{m.dataRequired}</p>
                  <div className="flex items-center gap-3 mt-1.5 text-xs text-[#6b7280] flex-wrap">
                    <span>Responsable: <span className="text-[#9ca3af]">{m.responsible}</span></span>
                    <span>Fuente: <span className="text-[#9ca3af]">{m.source}</span></span>
                    <span className="font-mono">{m.targetField}</span>
                  </div>
                  <div className="flex items-center gap-1.5 mt-1 text-xs text-red-400">
                    <CircleAlert size={11} />
                    <span>Si falta: {m.impactIfMissing}</span>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>

        {/* Gate de decisión ROI */}
        <div className="space-y-4">
          <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
            <h2 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2 mb-1">
              <Gauge size={15} className="text-[#f97316]" />
              Gate de escalamiento
            </h2>
            <p className="text-[#6b7280] text-xs mb-3">{ertPilot.roiPolicy}</p>
            {ertRoiGate.criteria.map(c => (
              <GateRow key={c.name} label={c.name} target={c.target} status={c.status} />
            ))}
          </div>

          <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
            <h2 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2 mb-3">
              <ShieldAlert size={15} className="text-[#f97316]" />
              Reglas de arranque del primer ciclo
            </h2>
            {ertRoiGate.startupRules.map(r => (
              <GateRow key={r.condition} label={r.condition} target={r.criterion} status={r.status} />
            ))}
          </div>

          <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
            <p className="text-[#9ca3af] text-xs leading-relaxed">{ertPilot.followUpRequest}</p>
          </div>
        </div>
      </div>

      {/* Roster de activos */}
      <div className="bg-[#111827] border border-[#1f2937] rounded-xl overflow-hidden">
        <div className="px-4 py-3 border-b border-[#1f2937] flex items-center justify-between">
          <h2 className="text-sm font-semibold text-[#f9fafb] flex items-center gap-2">
            <PackageSearch size={15} className="text-[#f97316]" />
            Maestro de activos — muestra piloto
          </h2>
          <span className="text-xs font-mono text-[#6b7280]">{ertAssets.length} activos · 02_ACTIVOS</span>
        </div>
        <div className="grid grid-cols-12 gap-2 px-4 py-2 border-b border-[#1f2937] text-xs text-[#6b7280] font-semibold uppercase tracking-wide">
          <div className="col-span-4">Asset_ID</div>
          <div className="col-span-2">Planta</div>
          <div className="col-span-3">Estado actual</div>
          <div className="col-span-3">Conciliado con ID físico</div>
        </div>
        <div className="divide-y divide-[#1f2937]">
          {visibleAssets.map(a => (
            <div key={a.id} className="grid grid-cols-12 gap-2 px-4 py-2.5 items-center">
              <div className="col-span-4 text-[#f97316] text-xs font-mono font-bold">{a.id}</div>
              <div className="col-span-2 text-[#9ca3af] text-xs">{a.plant}</div>
              <div className="col-span-3 text-[#9ca3af] text-xs">{a.status}</div>
              <div className="col-span-3 flex items-center gap-1.5 text-xs">
                {a.reconciled
                  ? <CircleCheck size={13} className="text-emerald-400" />
                  : <CircleAlert size={13} className="text-red-400" />}
                <span className={a.reconciled ? 'text-emerald-400' : 'text-red-400'}>
                  {a.reconciled ? 'Sí' : 'No'}
                </span>
              </div>
            </div>
          ))}
        </div>
        {ertAssets.length > 8 && (
          <button
            onClick={() => setShowAllAssets(v => !v)}
            className="w-full py-2.5 text-xs text-[#f97316] hover:bg-white/[0.02] border-t border-[#1f2937]"
          >
            {showAllAssets ? 'Mostrar menos' : `Ver los ${ertAssets.length} activos`}
          </button>
        )}
      </div>
    </div>
  )
}
