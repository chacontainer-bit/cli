import { useEffect } from 'react'
import { Outlet, useNavigate } from 'react-router-dom'
import { Bell, ShoppingCart } from 'lucide-react'
import Sidebar from './Sidebar'
import AssistantWidget from './AssistantWidget'
import { useApp } from '../context/AppContext'
import { useSimulator } from '../hooks/useSimulator'
import { roles } from '../data/mockData'

function TopBar() {
  const { alerts, cartCount, setCartOpen, role } = useApp()
  const criticalAlerts = alerts.filter(a => a.type === 'critico').length
  const roleData = roles.find(r => r.id === role)

  const now = new Date()
  const dateStr = now.toLocaleDateString('es-CL', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })

  return (
    <div className="h-14 bg-[#111827] border-b border-[#1f2937] flex items-center justify-between px-6 fixed top-0 left-60 right-0 z-30">
      <div>
        <p className="text-[#6b7280] text-xs capitalize">{dateStr}</p>
      </div>
      <div className="flex items-center gap-4">
        <button
          onClick={() => setCartOpen(true)}
          className="relative flex items-center gap-1.5 text-[#9ca3af] hover:text-[#f9fafb] transition-colors"
        >
          <ShoppingCart size={18} />
          {cartCount > 0 && (
            <span className="absolute -top-1.5 -right-1.5 bg-[#f97316] text-white text-xs w-4 h-4 rounded-full flex items-center justify-center font-bold">
              {cartCount}
            </span>
          )}
        </button>
        <div className="relative">
          <Bell size={18} className="text-[#9ca3af]" />
          {criticalAlerts > 0 && (
            <span className="absolute -top-1.5 -right-1.5 bg-[#ef4444] text-white text-xs w-4 h-4 rounded-full flex items-center justify-center font-bold animate-pulse">
              {criticalAlerts}
            </span>
          )}
        </div>
        {roleData && (
          <div
            className="w-7 h-7 rounded-full flex items-center justify-center text-xs font-bold text-white"
            style={{ backgroundColor: roleData.color }}
          >
            {roleData.avatar}
          </div>
        )}
      </div>
    </div>
  )
}

function SimulatorRunner() {
  useSimulator()
  return null
}

export default function Layout() {
  const { role } = useApp()
  const navigate = useNavigate()

  useEffect(() => {
    if (!role) navigate('/login')
  }, [role, navigate])

  if (!role) return null

  return (
    <div className="min-h-screen bg-[#0a0f1a]">
      <SimulatorRunner />
      <Sidebar />
      <TopBar />
      <main className="ml-60 pt-14 min-h-screen">
        <div className="p-6">
          <Outlet />
        </div>
      </main>
      <AssistantWidget />
    </div>
  )
}
