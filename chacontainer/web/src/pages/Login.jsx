import { useNavigate } from 'react-router-dom'
import { Package, ChevronRight, Shield, Truck, Wrench, BarChart3, Leaf, AlertTriangle } from 'lucide-react'
import { useApp } from '../context/AppContext'
import { roles } from '../data/mockData'

const roleIcons = {
  supply_chain: BarChart3,
  plant_manager: Wrench,
  logistics: Truck,
  packaging: Package,
  ehs: Leaf,
}

export default function Login() {
  const { setRole } = useApp()
  const navigate = useNavigate()

  function handleSelect(roleId) {
    setRole(roleId)
    navigate('/app/dashboard')
  }

  return (
    <div className="min-h-screen bg-[#0a0f1a] flex flex-col items-center justify-center p-8">
      <div className="fixed top-0 left-0 right-0 bg-amber-500/10 border-b border-amber-500/30 text-amber-400 text-xs font-semibold text-center py-2 flex items-center justify-center gap-1.5 z-50">
        <AlertTriangle size={13} />
        Ambiente de demostración — todos los datos (clientes, inventario, cifras) son de ejemplo, no reflejan la operación real
      </div>

      {/* Logo */}
      <div className="mb-12 text-center">
        <div className="flex items-center justify-center gap-3 mb-3">
          <div className="w-12 h-12 bg-[#f97316]/10 border border-[#f97316]/30 rounded-xl flex items-center justify-center">
            <Package className="text-[#f97316]" size={24} />
          </div>
          <h1 className="text-3xl font-bold tracking-widest text-[#f97316]">CHACONTAINER</h1>
        </div>
        <p className="text-[#6b7280] text-sm font-mono">Sistema ERP de Gestión de Contenedores Reutilizables</p>
        <div className="mt-3 flex items-center justify-center gap-2">
          <div className="h-px w-16 bg-[#1f2937]" />
          <span className="text-[#6b7280] text-xs">Seleccione su perfil para continuar</span>
          <div className="h-px w-16 bg-[#1f2937]" />
        </div>
      </div>

      {/* Role cards */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5 gap-4 max-w-5xl w-full">
        {roles.map(role => {
          const Icon = roleIcons[role.id] || Shield
          return (
            <button
              key={role.id}
              onClick={() => handleSelect(role.id)}
              className="group relative bg-[#111827] border border-[#1f2937] rounded-xl p-6 text-left transition-all duration-200 hover:border-opacity-100 hover:scale-[1.02] hover:bg-[#1a2235] focus:outline-none"
              style={{ '--role-color': role.color }}
            >
              <div
                className="absolute inset-0 rounded-xl opacity-0 group-hover:opacity-100 transition-opacity duration-200"
                style={{ boxShadow: `inset 0 0 0 1px ${role.color}40` }}
              />
              <div className="absolute top-0 left-0 right-0 h-0.5 rounded-t-xl opacity-60 group-hover:opacity-100 transition-opacity"
                style={{ backgroundColor: role.color }} />

              <div
                className="w-12 h-12 rounded-xl flex items-center justify-center mb-4 transition-all duration-200"
                style={{ backgroundColor: `${role.color}20`, border: `1px solid ${role.color}40` }}
              >
                <Icon size={22} style={{ color: role.color }} />
              </div>

              <div
                className="text-xs font-bold font-mono mb-1 transition-colors"
                style={{ color: role.color }}
              >
                {role.avatar}
              </div>
              <h3 className="text-[#f9fafb] font-semibold text-sm mb-2 leading-tight">{role.name}</h3>
              <p className="text-[#6b7280] text-xs leading-relaxed">{role.description}</p>

              <div className="mt-4 flex items-center gap-1 text-xs font-medium opacity-0 group-hover:opacity-100 transition-opacity"
                style={{ color: role.color }}>
                <span>Ingresar</span>
                <ChevronRight size={12} />
              </div>
            </button>
          )
        })}
      </div>

      <div className="mt-12 text-center">
        <p className="text-[#374151] text-xs font-mono">v2.4.1 · Chacontainer S.A. de C.V. · Querétaro, México</p>
      </div>
    </div>
  )
}
