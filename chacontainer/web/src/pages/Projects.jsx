import { useState } from 'react'
import { X, Users, Calendar, DollarSign, Package, AlertTriangle, ChevronRight } from 'lucide-react'
import { projects } from '../data/mockData'

const columns = [
  { id: 'propuesta', label: 'Propuesta', color: '#6b7280' },
  { id: 'en_curso', label: 'En Curso', color: '#3b82f6' },
  { id: 'cerrado', label: 'Cerrado', color: '#10b981' },
]

const riskConfig = {
  bajo: { label: 'Riesgo Bajo', bg: 'bg-emerald-900/40', text: 'text-emerald-400', border: 'border-emerald-700' },
  medio: { label: 'Riesgo Medio', bg: 'bg-amber-900/40', text: 'text-amber-400', border: 'border-amber-700' },
  alto: { label: 'Riesgo Alto', bg: 'bg-red-900/40', text: 'text-red-400', border: 'border-red-700' },
}

function phaseColor(phase, total) {
  const ratio = phase / total
  if (ratio >= 0.8) return 'bg-emerald-500'
  if (ratio >= 0.4) return 'bg-amber-500'
  return 'bg-blue-500'
}

function ProjectCard({ project, onClick }) {
  const risk = riskConfig[project.risk]
  const burnPct = project.budget > 0 ? Math.round((project.spent / project.budget) * 100) : 0
  const phasePct = Math.round((project.phase / project.totalPhases) * 100)

  return (
    <div
      onClick={() => onClick(project)}
      className="bg-[#111827] border border-[#1f2937] rounded-xl p-4 cursor-pointer hover:border-[#374151] hover:bg-[#1a2235] transition-all duration-150 space-y-3"
    >
      <div className="flex items-start justify-between gap-2">
        <div className="flex-1 min-w-0">
          <p className="text-[#6b7280] text-xs font-mono">{project.id}</p>
          <h3 className="text-[#f9fafb] text-sm font-semibold mt-0.5 leading-tight">{project.name}</h3>
        </div>
        <ChevronRight size={14} className="text-[#374151] flex-shrink-0 mt-1" />
      </div>

      <div className="flex items-center gap-2 flex-wrap">
        <span className="bg-[#0a0f1a] border border-[#374151] text-[#9ca3af] text-xs px-2 py-0.5 rounded font-medium">
          {project.client}
        </span>
        <span className={`text-xs px-2 py-0.5 rounded border ${risk.bg} ${risk.text} ${risk.border}`}>
          {risk.label}
        </span>
      </div>

      <div>
        <div className="flex justify-between text-xs mb-1">
          <span className="text-[#6b7280]">Fase {project.phase}/{project.totalPhases}</span>
          <span className="text-[#9ca3af] font-mono">{phasePct}%</span>
        </div>
        <div className="h-1.5 bg-[#1f2937] rounded-full overflow-hidden">
          <div
            className={`h-full rounded-full transition-all ${phaseColor(project.phase, project.totalPhases)}`}
            style={{ width: `${phasePct}%` }}
          />
        </div>
      </div>

      <div className="grid grid-cols-2 gap-2 text-xs">
        <div>
          <p className="text-[#6b7280]">Presupuesto</p>
          <p className="text-[#f9fafb] font-mono">${(project.budget / 1000000).toFixed(1)}M</p>
        </div>
        <div>
          <p className="text-[#6b7280]">Ejecutado</p>
          <p className={`font-mono ${burnPct > 90 ? 'text-red-400' : burnPct > 70 ? 'text-amber-400' : 'text-[#f9fafb]'}`}>
            {burnPct}%
          </p>
        </div>
      </div>

      <div className="flex items-center justify-between text-xs">
        <div className="flex items-center gap-1.5 text-[#6b7280]">
          <Users size={11} />
          <span>{project.manager}</span>
        </div>
        <div className="flex items-center gap-1.5 text-[#6b7280]">
          <Package size={11} />
          <span>{project.containers} cont.</span>
        </div>
      </div>
    </div>
  )
}

function ProjectDetail({ project, onClose }) {
  const risk = riskConfig[project.risk]
  const burnPct = project.budget > 0 ? Math.round((project.spent / project.budget) * 100) : 0
  const phasePct = Math.round((project.phase / project.totalPhases) * 100)

  const phases = Array.from({ length: project.totalPhases }, (_, i) => ({
    num: i + 1,
    label: [`Diagnóstico`, `Propuesta`, `Aprobación`, `Implementación`, `Cierre`, `Expansión`][i] || `Fase ${i + 1}`,
    done: i + 1 <= project.phase,
  }))

  return (
    <>
      <div className="fixed inset-0 bg-black/50 z-40" onClick={onClose} />
      <div className="fixed right-0 top-0 h-full w-full max-w-lg bg-[#111827] border-l border-[#1f2937] z-50 flex flex-col overflow-y-auto">
        <div className="p-5 border-b border-[#1f2937] flex items-center justify-between sticky top-0 bg-[#111827]">
          <div>
            <p className="text-xs text-[#6b7280] font-mono">{project.id}</p>
            <h2 className="text-[#f9fafb] font-bold text-sm mt-0.5">{project.name}</h2>
          </div>
          <button onClick={onClose} className="text-[#6b7280] hover:text-[#f9fafb]">
            <X size={18} />
          </button>
        </div>

        <div className="p-5 space-y-5">
          {/* Status badges */}
          <div className="flex gap-2 flex-wrap">
            <span className="bg-[#0a0f1a] border border-[#374151] text-[#9ca3af] text-xs px-3 py-1 rounded-full font-medium">
              {project.client}
            </span>
            <span className={`text-xs px-3 py-1 rounded-full border ${risk.bg} ${risk.text} ${risk.border}`}>
              {risk.label}
            </span>
            {project.risk === 'alto' && (
              <div className="flex items-center gap-1 text-red-400 text-xs">
                <AlertTriangle size={12} />
                <span>Requiere atención</span>
              </div>
            )}
          </div>

          {/* Phase timeline */}
          <div>
            <p className="text-xs text-[#6b7280] font-semibold uppercase tracking-wide mb-3">Fases del Proyecto</p>
            <div className="flex gap-1">
              {phases.map(p => (
                <div key={p.num} className="flex-1">
                  <div className={`h-2 rounded-full mb-1 ${p.done ? phaseColor(project.phase, project.totalPhases) : 'bg-[#1f2937]'}`} />
                  <p className="text-[#6b7280] text-xs truncate">{p.label}</p>
                </div>
              ))}
            </div>
            <p className="text-right text-xs text-[#9ca3af] font-mono mt-1">Fase {project.phase} de {project.totalPhases} ({phasePct}%)</p>
          </div>

          {/* Dates */}
          <div className="grid grid-cols-2 gap-4">
            <div className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg p-3">
              <div className="flex items-center gap-1.5 text-[#6b7280] text-xs mb-1">
                <Calendar size={12} />
                <span>Inicio</span>
              </div>
              <p className="text-[#f9fafb] text-sm font-mono">{project.startDate}</p>
            </div>
            <div className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg p-3">
              <div className="flex items-center gap-1.5 text-[#6b7280] text-xs mb-1">
                <Calendar size={12} />
                <span>Término</span>
              </div>
              <p className="text-[#f9fafb] text-sm font-mono">{project.endDate}</p>
            </div>
          </div>

          {/* Budget */}
          <div className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg p-4">
            <div className="flex items-center gap-1.5 text-[#6b7280] text-xs mb-3">
              <DollarSign size={12} />
              <span>Presupuesto</span>
            </div>
            <div className="flex justify-between mb-2">
              <div>
                <p className="text-xs text-[#6b7280]">Presupuesto Total</p>
                <p className="text-[#f9fafb] font-mono font-bold">${project.budget.toLocaleString('es-CL')}</p>
              </div>
              <div className="text-right">
                <p className="text-xs text-[#6b7280]">Ejecutado</p>
                <p className="text-[#f97316] font-mono font-bold">${project.spent.toLocaleString('es-CL')}</p>
              </div>
            </div>
            <div className="h-2 bg-[#1f2937] rounded-full overflow-hidden">
              <div
                className={`h-full rounded-full ${burnPct > 90 ? 'bg-red-500' : burnPct > 70 ? 'bg-amber-500' : 'bg-emerald-500'}`}
                style={{ width: `${burnPct}%` }}
              />
            </div>
            <p className="text-right text-xs text-[#6b7280] mt-1 font-mono">{burnPct}% ejecutado</p>
          </div>

          {/* Stats */}
          <div className="grid grid-cols-3 gap-3">
            {[
              { label: 'Contenedores', value: project.containers, icon: Package },
              { label: 'Manager', value: project.manager.split(' ')[0], icon: Users },
              { label: 'Fases Total', value: project.totalPhases, icon: ChevronRight },
            ].map(({ label, value, icon: Icon }) => (
              <div key={label} className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg p-3 text-center">
                <Icon size={16} className="text-[#f97316] mx-auto mb-1" />
                <p className="text-[#f9fafb] font-bold text-sm">{value}</p>
                <p className="text-[#6b7280] text-xs">{label}</p>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  )
}

export default function Projects() {
  const [selectedProject, setSelectedProject] = useState(null)

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-xl font-bold text-[#f9fafb]">Proyectos</h1>
          <p className="text-[#6b7280] text-sm">Tablero Kanban de proyectos de implementación</p>
        </div>
        <div className="flex gap-2 text-xs text-[#6b7280]">
          <span className="font-mono">{projects.length} proyectos</span>
          <span>·</span>
          <span className="text-blue-400">{projects.filter(p => p.status === 'en_curso').length} activos</span>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        {columns.map(col => {
          const colProjects = projects.filter(p => p.status === col.id)
          return (
            <div key={col.id} className="bg-[#111827] border border-[#1f2937] rounded-xl">
              <div className="px-4 py-3 border-b border-[#1f2937] flex items-center gap-2">
                <div className="w-2.5 h-2.5 rounded-full" style={{ backgroundColor: col.color }} />
                <h2 className="text-sm font-semibold text-[#f9fafb]">{col.label}</h2>
                <span className="ml-auto text-xs font-mono text-[#6b7280] bg-[#0a0f1a] px-2 py-0.5 rounded">
                  {colProjects.length}
                </span>
              </div>
              <div className="p-3 space-y-3 min-h-64">
                {colProjects.map(p => (
                  <ProjectCard key={p.id} project={p} onClick={setSelectedProject} />
                ))}
                {colProjects.length === 0 && (
                  <div className="flex items-center justify-center h-32 text-[#374151] text-sm">
                    Sin proyectos
                  </div>
                )}
              </div>
            </div>
          )
        })}
      </div>

      {selectedProject && (
        <ProjectDetail project={selectedProject} onClose={() => setSelectedProject(null)} />
      )}
    </div>
  )
}
