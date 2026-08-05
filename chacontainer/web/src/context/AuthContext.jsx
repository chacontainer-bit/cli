import { createContext, useContext, useState, useCallback, useEffect } from 'react'
import { api, getSession, saveSession, clearSession } from '../lib/api'

const AuthContext = createContext(null)

export function AuthProvider({ children }) {
  const [session, setSession] = useState(() => getSession())
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  useEffect(() => {
    function handleUnauthorized() {
      setSession(null)
    }
    window.addEventListener('chacontainer:unauthorized', handleUnauthorized)
    return () => window.removeEventListener('chacontainer:unauthorized', handleUnauthorized)
  }, [])

  const login = useCallback(async (email, password) => {
    setLoading(true)
    setError(null)
    try {
      const res = await api.login({ email, password })
      saveSession(res)
      setSession({ user: res.user, tenant: res.tenant })
      return res
    } catch (err) {
      setError(err.message)
      throw err
    } finally {
      setLoading(false)
    }
  }, [])

  const register = useCallback(async (companyName, name, email, password) => {
    setLoading(true)
    setError(null)
    try {
      const res = await api.register({ company_name: companyName, name, email, password })
      saveSession(res)
      setSession({ user: res.user, tenant: res.tenant })
      return res
    } catch (err) {
      setError(err.message)
      throw err
    } finally {
      setLoading(false)
    }
  }, [])

  const logout = useCallback(() => {
    clearSession()
    setSession(null)
  }, [])

  return (
    <AuthContext.Provider value={{
      user: session?.user || null,
      tenant: session?.tenant || null,
      isAuthenticated: !!session?.user,
      loading, error, setError,
      login, register, logout,
    }}>
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth() {
  const ctx = useContext(AuthContext)
  if (!ctx) throw new Error('useAuth must be used within AuthProvider')
  return ctx
}
