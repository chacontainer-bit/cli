import { createContext, useContext, useState, useCallback } from 'react'
import { recentActivity, alerts as initialAlerts } from '../data/mockData'
import { getSession, setSession as persistSession, clearSession as clearPersistedSession } from '../lib/api'

const AppContext = createContext(null)

export function AppProvider({ children }) {
  const [role, setRole] = useState(null)
  // `session` is only set when logged in against the real local backend
  // (see Login.jsx "Modo local"). The demo role picker never touches this,
  // so the existing mock-data experience is unaffected either way.
  const [session, setSessionState] = useState(() => getSession())

  const loginWithSession = useCallback((session) => {
    persistSession(session)
    setSessionState(session)
  }, [])

  const logout = useCallback(() => {
    clearPersistedSession()
    setSessionState(null)
    setRole(null)
  }, [])
  const [cart, setCart] = useState([])
  const [cartOpen, setCartOpen] = useState(false)
  const [activity, setActivity] = useState(recentActivity)
  const [alerts, setAlerts] = useState(initialAlerts)
  const [selectedWarehouse, setSelectedWarehouse] = useState('bodega_norte')

  const addToCart = useCallback((product, qty) => {
    setCart(prev => {
      const existing = prev.find(i => i.id === product.id)
      if (existing) {
        return prev.map(i => i.id === product.id ? { ...i, qty: i.qty + qty } : i)
      }
      return [...prev, { ...product, qty }]
    })
  }, [])

  const removeFromCart = useCallback((productId) => {
    setCart(prev => prev.filter(i => i.id !== productId))
  }, [])

  const updateCartQty = useCallback((productId, qty) => {
    if (qty <= 0) {
      removeFromCart(productId)
      return
    }
    setCart(prev => prev.map(i => i.id === productId ? { ...i, qty } : i))
  }, [removeFromCart])

  const clearCart = useCallback(() => setCart([]), [])

  const addActivity = useCallback((item) => {
    setActivity(prev => [item, ...prev].slice(0, 20))
  }, [])

  const addAlert = useCallback((alert) => {
    setAlerts(prev => [alert, ...prev].slice(0, 10))
  }, [])

  const cartTotal = cart.reduce((sum, i) => sum + i.price * i.qty, 0)
  const cartCount = cart.reduce((sum, i) => sum + i.qty, 0)

  return (
    <AppContext.Provider value={{
      role, setRole,
      session, loginWithSession, logout,
      cart, cartOpen, setCartOpen,
      addToCart, removeFromCart, updateCartQty, clearCart,
      cartTotal, cartCount,
      activity, addActivity,
      alerts, addAlert,
      selectedWarehouse, setSelectedWarehouse,
    }}>
      {children}
    </AppContext.Provider>
  )
}

export function useApp() {
  const ctx = useContext(AppContext)
  if (!ctx) throw new Error('useApp must be used within AppProvider')
  return ctx
}
