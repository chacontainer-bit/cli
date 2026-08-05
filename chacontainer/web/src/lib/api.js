const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'
const TOKEN_KEY = 'chacontainer_token'
const SESSION_KEY = 'chacontainer_session'

export function getToken() {
  return localStorage.getItem(TOKEN_KEY) || ''
}

export function getSession() {
  const raw = localStorage.getItem(SESSION_KEY)
  if (!raw) return null
  try {
    return JSON.parse(raw)
  } catch {
    return null
  }
}

export function saveSession({ token, user, tenant }) {
  localStorage.setItem(TOKEN_KEY, token)
  localStorage.setItem(SESSION_KEY, JSON.stringify({ user, tenant }))
}

export function clearSession() {
  localStorage.removeItem(TOKEN_KEY)
  localStorage.removeItem(SESSION_KEY)
}

export class ApiError extends Error {
  constructor(message, status) {
    super(message)
    this.status = status
  }
}

async function request(path, { method = 'GET', body, auth = true } = {}) {
  const headers = { 'Content-Type': 'application/json' }
  if (auth) {
    const token = getToken()
    if (token) headers.Authorization = `Bearer ${token}`
  }

  const res = await fetch(`${BASE_URL}${path}`, {
    method,
    headers,
    body: body !== undefined ? JSON.stringify(body) : undefined,
  })

  if (res.status === 401 && auth) {
    clearSession()
    window.dispatchEvent(new CustomEvent('chacontainer:unauthorized'))
  }

  let data = null
  const text = await res.text()
  if (text) {
    try {
      data = JSON.parse(text)
    } catch {
      data = null
    }
  }

  if (!res.ok) {
    const message = data?.error || `Error ${res.status}`
    throw new ApiError(message, res.status)
  }
  return data
}

export const api = {
  register: (payload) => request('/api/v1/auth/register', { method: 'POST', body: payload, auth: false }),
  login: (payload) => request('/api/v1/auth/login', { method: 'POST', body: payload, auth: false }),

  dashboardSummary: () => request('/api/v1/dashboard/summary'),
  assetTrend: (days = 30) => request(`/api/v1/dashboard/assets/trend?days=${days}`),
  shipmentTrend: (days = 30) => request(`/api/v1/dashboard/shipments/trend?days=${days}`),

  listAssets: (params = {}) => request(`/api/v1/assets?${new URLSearchParams(params)}`),
  createAsset: (payload) => request('/api/v1/assets', { method: 'POST', body: payload }),
  getAssetEvents: (id) => request(`/api/v1/assets/${id}/events`),
  scanAsset: (payload) => request('/api/v1/assets/scan', { method: 'POST', body: payload }),

  listShipments: (params = {}) => request(`/api/v1/shipments?${new URLSearchParams(params)}`),
  createShipment: (payload) => request('/api/v1/shipments', { method: 'POST', body: payload }),
  transitionShipment: (id, payload) => request(`/api/v1/shipments/${id}/transition`, { method: 'POST', body: payload }),
  getShipmentTimeline: (id) => request(`/api/v1/shipments/${id}/timeline`),

  listClients: (params = {}) => request(`/api/v1/clients?${new URLSearchParams(params)}`),
  createClient: (payload) => request('/api/v1/clients', { method: 'POST', body: payload }),

  listPlants: () => request('/api/v1/plants'),
  createPlant: (payload) => request('/api/v1/plants', { method: 'POST', body: payload }),
  listZones: (plantId) => request(`/api/v1/plants/${plantId}/zones`),
  createZone: (plantId, payload) => request(`/api/v1/plants/${plantId}/zones`, { method: 'POST', body: payload }),
}
