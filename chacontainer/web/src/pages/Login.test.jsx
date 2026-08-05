import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { MemoryRouter, Routes, Route } from 'react-router-dom'
import Login from './Login'
import { AuthProvider } from '../context/AuthContext'

function renderLogin() {
  return render(
    <MemoryRouter initialEntries={['/login']}>
      <AuthProvider>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/dashboard" element={<div>Dashboard Content</div>} />
        </Routes>
      </AuthProvider>
    </MemoryRouter>
  )
}

describe('Login', () => {
  beforeEach(() => {
    localStorage.clear()
    globalThis.fetch = vi.fn()
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('shows an error and stays on the login page for invalid credentials', async () => {
    globalThis.fetch.mockResolvedValueOnce({
      ok: false,
      status: 401,
      text: async () => JSON.stringify({ error: 'invalid credentials' }),
    })

    renderLogin()
    const user = userEvent.setup()

    await user.type(screen.getByPlaceholderText('tu@empresa.mx'), 'demo@chacontainer.mx')
    await user.type(screen.getByPlaceholderText('••••••••'), 'wrong-password')
    await user.click(screen.getByRole('button', { name: 'Ingresar' }))

    expect(await screen.findByText('invalid credentials')).toBeInTheDocument()
    expect(screen.queryByText('Dashboard Content')).not.toBeInTheDocument()
  })

  it('navigates to the dashboard and persists the session on successful login', async () => {
    globalThis.fetch.mockResolvedValueOnce({
      ok: true,
      status: 200,
      text: async () => JSON.stringify({
        token: 'fake-jwt-token',
        user: { id: 'u1', name: 'Admin Demo', role: 'admin' },
        tenant: { id: 't1', name: 'CHACONTAINER Demo' },
      }),
    })

    renderLogin()
    const user = userEvent.setup()

    await user.type(screen.getByPlaceholderText('tu@empresa.mx'), 'demo@chacontainer.mx')
    await user.type(screen.getByPlaceholderText('••••••••'), 'Demo12345!')
    await user.click(screen.getByRole('button', { name: 'Ingresar' }))

    expect(await screen.findByText('Dashboard Content')).toBeInTheDocument()
    expect(localStorage.getItem('chacontainer_token')).toBe('fake-jwt-token')
  })

  it('switches to the register form, revealing company name and full name fields', async () => {
    renderLogin()
    const user = userEvent.setup()

    await user.click(screen.getByRole('button', { name: 'Crear cuenta' }))
    expect(screen.getByPlaceholderText('Autopartes del Bajío')).toBeInTheDocument()
    expect(screen.getByPlaceholderText('Ana Ramírez')).toBeInTheDocument()
  })
})
