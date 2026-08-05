import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen } from '@testing-library/react'
import Dashboard from './Dashboard'
import { AuthProvider } from '../context/AuthContext'

vi.mock('../lib/api', async (importOriginal) => {
  const actual = await importOriginal()
  return {
    ...actual,
    api: {
      ...actual.api,
      dashboardSummary: vi.fn(),
      assetTrend: vi.fn(),
      shipmentTrend: vi.fn(),
    },
  }
})

import { api } from '../lib/api'

function seedSession() {
  localStorage.setItem('chacontainer_token', 'fake-token')
  localStorage.setItem('chacontainer_session', JSON.stringify({
    user: { id: 'u1', name: 'Admin Demo', role: 'admin' },
    tenant: { id: 't1', name: 'CHACONTAINER Demo' },
  }))
}

describe('Dashboard', () => {
  beforeEach(() => {
    localStorage.clear()
    seedSession()
  })

  it('renders KPI values from the dashboard summary API', async () => {
    api.dashboardSummary.mockResolvedValue({
      assets_by_status: { available: 2, in_transit: 1 },
      assets_by_type: {},
      shipments_by_status: {},
      plant_occupancy: [{ plant_id: 'p1', plant_name: 'Planta Norte', total_slots: 100, used_slots: 30, occupancy_pct: 30 }],
      late_shipments: 1,
      total_assets: 3,
      assets_in_transit: 1,
      shipments_today: 2,
      scan_events_24h: 5,
      top_clients: [],
    })
    api.assetTrend.mockResolvedValue({ data: [{ date: '2026-08-01', value: 1 }] })
    api.shipmentTrend.mockResolvedValue({ data: [{ date: '2026-08-01', value: 1 }] })

    render(<AuthProvider><Dashboard /></AuthProvider>)

    expect(await screen.findByText('3')).toBeInTheDocument() // total_assets
    expect(screen.getByText('Admin Demo')).toBeInTheDocument()
    expect(screen.getByText('Planta Norte')).toBeInTheDocument()
    expect(screen.getByText('30/100')).toBeInTheDocument()
  })

  it('shows an error message when the summary request fails', async () => {
    api.dashboardSummary.mockRejectedValue(new Error('boom'))
    api.assetTrend.mockResolvedValue({ data: [] })
    api.shipmentTrend.mockResolvedValue({ data: [] })

    render(<AuthProvider><Dashboard /></AuthProvider>)

    expect(await screen.findByText(/No se pudo cargar el dashboard/)).toBeInTheDocument()
  })
})
