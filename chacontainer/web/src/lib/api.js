// Thin client for the real CHACONTAINER backend. Everything here talks to
// your own local API (same origin via the nginx proxy in docker-compose, or
// VITE_API_URL in dev) - there is no call to any third-party service.
const API_URL = import.meta.env.VITE_API_URL || ''
const STORAGE_KEY = 'chacontainer_session'

export function getSession() {
  try {
    const raw = localStorage.getItem(STORAGE_KEY)
    return raw ? JSON.parse(raw) : null
  } catch {
    return null
  }
}

export function setSession(session) {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(session))
}

export function clearSession() {
  localStorage.removeItem(STORAGE_KEY)
}

async function request(path, options = {}) {
  const session = getSession()
  const headers = { 'Content-Type': 'application/json', ...(options.headers || {}) }
  if (session?.token) headers.Authorization = `Bearer ${session.token}`

  const res = await fetch(`${API_URL}${path}`, { ...options, headers })
  if (!res.ok) {
    let message = `HTTP ${res.status}`
    try {
      const body = await res.json()
      if (body?.error) message = body.error
    } catch {
      // response had no JSON body; keep the generic message
    }
    throw new Error(message)
  }
  if (res.status === 204) return null
  return res.json()
}

export const api = {
  login: (email, password) =>
    request('/api/v1/auth/login', { method: 'POST', body: JSON.stringify({ email, password }) }),
  dashboardSummary: () => request('/api/v1/dashboard/summary'),
  listAssets: (query = '') => request(`/api/v1/assets${query}`),
  listPlants: () => request('/api/v1/plants'),
  listClients: (query = '') => request(`/api/v1/clients${query}`),
  listShipments: (query = '') => request(`/api/v1/shipments${query}`),
}
