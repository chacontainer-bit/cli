import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Package, Loader2 } from 'lucide-react'
import { useAuth } from '../context/AuthContext'

export default function Login() {
  const { login, register, loading, error, setError } = useAuth()
  const navigate = useNavigate()
  const [mode, setMode] = useState('login') // 'login' | 'register'
  const [form, setForm] = useState({ companyName: '', name: '', email: '', password: '' })

  function update(field) {
    return (e) => setForm((f) => ({ ...f, [field]: e.target.value }))
  }

  async function handleSubmit(e) {
    e.preventDefault()
    setError(null)
    try {
      if (mode === 'login') {
        await login(form.email, form.password)
      } else {
        await register(form.companyName, form.name, form.email, form.password)
      }
      navigate('/dashboard')
    } catch {
      // error is surfaced via useAuth().error
    }
  }

  return (
    <div className="min-h-screen bg-[#0a0f1a] flex flex-col items-center justify-center p-8">
      <div className="mb-8 text-center">
        <div className="flex items-center justify-center gap-3 mb-3">
          <div className="w-12 h-12 bg-[#f97316]/10 border border-[#f97316]/30 rounded-xl flex items-center justify-center">
            <Package className="text-[#f97316]" size={24} />
          </div>
          <h1 className="text-3xl font-bold tracking-widest text-[#f97316]">CHACONTAINER</h1>
        </div>
        <p className="text-[#6b7280] text-sm font-mono">Gestión de activos retornables multi-planta</p>
      </div>

      <div className="w-full max-w-sm bg-[#111827] border border-[#1f2937] rounded-xl p-6">
        <div className="flex gap-1 mb-6 bg-[#0a0f1a] rounded-lg p-1">
          <button
            type="button"
            onClick={() => { setMode('login'); setError(null) }}
            className={`flex-1 text-sm py-2 rounded-md transition-colors ${mode === 'login' ? 'bg-[#f97316] text-white' : 'text-[#9ca3af]'}`}
          >
            Iniciar sesión
          </button>
          <button
            type="button"
            onClick={() => { setMode('register'); setError(null) }}
            className={`flex-1 text-sm py-2 rounded-md transition-colors ${mode === 'register' ? 'bg-[#f97316] text-white' : 'text-[#9ca3af]'}`}
          >
            Crear cuenta
          </button>
        </div>

        <form onSubmit={handleSubmit} className="space-y-3">
          {mode === 'register' && (
            <>
              <div>
                <label className="text-[#9ca3af] text-xs mb-1 block">Nombre de la empresa</label>
                <input
                  required
                  value={form.companyName}
                  onChange={update('companyName')}
                  placeholder="Autopartes del Bajío"
                  className="w-full bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]"
                />
              </div>
              <div>
                <label className="text-[#9ca3af] text-xs mb-1 block">Tu nombre</label>
                <input
                  required
                  value={form.name}
                  onChange={update('name')}
                  placeholder="Ana Ramírez"
                  className="w-full bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]"
                />
              </div>
            </>
          )}
          <div>
            <label className="text-[#9ca3af] text-xs mb-1 block">Email</label>
            <input
              required
              type="email"
              value={form.email}
              onChange={update('email')}
              placeholder="tu@empresa.mx"
              className="w-full bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]"
            />
          </div>
          <div>
            <label className="text-[#9ca3af] text-xs mb-1 block">Contraseña</label>
            <input
              required
              type="password"
              minLength={8}
              value={form.password}
              onChange={update('password')}
              placeholder="••••••••"
              className="w-full bg-[#0a0f1a] border border-[#1f2937] rounded-lg px-3 py-2 text-sm text-[#f9fafb] placeholder-[#4b5563] focus:outline-none focus:border-[#f97316]"
            />
            {mode === 'register' && (
              <p className="text-[#4b5563] text-xs mt-1">Mínimo 8 caracteres</p>
            )}
          </div>

          {error && (
            <div className="text-red-400 text-xs bg-red-500/10 border border-red-500/30 rounded-lg px-3 py-2">
              {error}
            </div>
          )}

          <button
            type="submit"
            disabled={loading}
            className="w-full bg-[#f97316] hover:bg-[#ea6a0d] disabled:opacity-60 text-white text-sm font-semibold py-2.5 rounded-lg transition-colors flex items-center justify-center gap-2"
          >
            {loading && <Loader2 size={14} className="animate-spin" />}
            {mode === 'login' ? 'Ingresar' : 'Crear cuenta'}
          </button>
        </form>
      </div>

      <div className="mt-8 text-center">
        <p className="text-[#374151] text-xs font-mono">Chacontainer · Querétaro, México</p>
      </div>
    </div>
  )
}
