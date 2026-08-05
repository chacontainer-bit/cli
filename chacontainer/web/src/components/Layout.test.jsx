import { describe, it, expect, beforeEach } from 'vitest'
import { render, screen } from '@testing-library/react'
import { MemoryRouter, Routes, Route } from 'react-router-dom'
import Layout from './Layout'
import { AuthProvider } from '../context/AuthContext'
import { AppProvider } from '../context/AppContext'

function renderApp(initialPath) {
  return render(
    <MemoryRouter initialEntries={[initialPath]}>
      <AuthProvider>
        <AppProvider>
          <Routes>
            <Route path="/login" element={<div>Login Page</div>} />
            <Route path="/" element={<Layout />}>
              <Route path="dashboard" element={<div>Protected Dashboard</div>} />
            </Route>
          </Routes>
        </AppProvider>
      </AuthProvider>
    </MemoryRouter>
  )
}

describe('Layout route protection', () => {
  beforeEach(() => {
    localStorage.clear()
  })

  it('redirects to /login when there is no active session', async () => {
    renderApp('/dashboard')
    expect(await screen.findByText('Login Page')).toBeInTheDocument()
    expect(screen.queryByText('Protected Dashboard')).not.toBeInTheDocument()
  })

  it('renders the protected content when a session exists', async () => {
    localStorage.setItem('chacontainer_token', 'fake-token')
    localStorage.setItem('chacontainer_session', JSON.stringify({
      user: { id: 'u1', name: 'Admin Demo', role: 'admin' },
      tenant: { id: 't1', name: 'CHACONTAINER Demo' },
    }))

    renderApp('/dashboard')
    expect(await screen.findByText('Protected Dashboard')).toBeInTheDocument()
  })
})
