import { NavLink, useNavigate } from 'react-router-dom'
import {
  LayoutDashboard, ShoppingCart, FolderKanban, MapPin,
  Calculator, Settings, LogOut, Package, ChevronRight, Bot, Layers
} from 'lucide-react'
import { useApp } from '../context/AppContext'
import { roles } from '../data/mockData'

const navItems = [
  { to: '/dashboard', icon: LayoutDashboard, label: 'Dashboard' },
  { to: '/marketplace', icon: ShoppingCart, label: 'Marketplace' },
  { to: '/projects', icon: FolderKanban, label: 'Proyectos' },
  { to: '/traceability', icon: MapPin, label: 'Trazabilidad' },
  { to: '/calculator', icon: Calculator, label: 'Calculadora' },
  { to: '/backoffice', icon: Settings, label: 'Backoffice' },
  { to: '/asistente', icon: Bot, label: 'Asistente' },
  { to: '/assets', icon: Layers, label: 'Fábrica de Assets' },
]

export default function Sidebar() {
  const { role, setRole } = useApp()
  const navigate = useNavigate()
  const roleData = roles.find(r => r.id === role)

  function handleLogout() {
    setRole(null)
    navigate('/login')
  }

  return (
    <aside className="fixed left-0 top-0 h-screen w-60 bg-[#111827] border-r border-[#1f2937] flex flex-col z-40">
      {/* Logo */}
      <div className="px-5 py-5 border-b border-[#1f2937]">
        <div className="flex items-center gap-2">
          <Package className="text-[#f97316]" size={22} />
          <span className="text-[#f97316] font-bold tracking-widest text-sm">CHACONTAINER</span>
        </div>
        <p className="text-[#6b7280] text-xs mt-1 font-mono">ERP v2.4.1</p>
      </div>

      {/* Nav */}
      <nav className="flex-1 py-4 overflow-y-auto">
        {navItems.map(({ to, icon: Icon, label }) => (
          <NavLink
            key={to}
            to={to}
            className={({ isActive }) =>
              `flex items-center gap-3 px-5 py-3 text-sm transition-all duration-150 group ${
                isActive
                  ? 'text-[#f97316] bg-[#f97316]/10 border-r-2 border-[#f97316]'
                  : 'text-[#9ca3af] hover:text-[#f9fafb] hover:bg-white/5'
              }`
            }
          >
            {({ isActive }) => (
              <>
                <Icon size={18} className={isActive ? 'text-[#f97316]' : 'text-[#6b7280] group-hover:text-[#9ca3af]'} />
                <span className="font-medium">{label}</span>
                {isActive && <ChevronRight size={14} className="ml-auto text-[#f97316]" />}
              </>
            )}
          </NavLink>
        ))}
      </nav>

      {/* Role badge + logout */}
      <div className="border-t border-[#1f2937] p-4">
        {roleData && (
          <div className="flex items-center gap-3 mb-3">
            <div
              className="w-8 h-8 rounded-full flex items-center justify-center text-xs font-bold text-white flex-shrink-0"
              style={{ backgroundColor: roleData.color }}
            >
              {roleData.avatar}
            </div>
            <div className="min-w-0">
              <p className="text-[#f9fafb] text-xs font-medium truncate">{roleData.name}</p>
              <p className="text-[#6b7280] text-xs truncate font-mono">{roleData.id}</p>
            </div>
          </div>
        )}
        <button
          onClick={handleLogout}
          className="flex items-center gap-2 text-[#6b7280] hover:text-[#ef4444] text-xs transition-colors w-full py-1"
        >
          <LogOut size={14} />
          <span>Cerrar sesión</span>
        </button>
      </div>
    </aside>
  )
}
