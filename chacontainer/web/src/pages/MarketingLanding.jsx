import { useNavigate } from 'react-router-dom'
import { Package, Rocket, Eye } from 'lucide-react'

export default function MarketingLanding() {
  const navigate = useNavigate()

  return (
    <div className="min-h-screen bg-[#0a0f1a]">
      {/* Nav */}
      <header className="bg-[#0d1220] border-b border-[#1f2937] px-8 py-4 flex items-center justify-between">
        <div className="flex items-center gap-3">
          <div className="w-11 h-11 rounded-xl border-2 border-[#f97316] flex items-center justify-center">
            <Package className="text-[#f97316]" size={22} />
          </div>
          <div>
            <p className="text-white font-bold tracking-wide text-lg leading-tight">CHACONTAINER</p>
            <p className="text-[#6b7280] text-xs leading-tight">Plataforma de gobernanza de envases</p>
          </div>
        </div>

        <nav className="hidden md:flex items-center gap-8 text-sm text-[#d1d5db]">
          <a href="#problema" className="hover:text-white transition-colors">Problema</a>
          <a href="#soluciones" className="hover:text-white transition-colors">Soluciones</a>
          <a href="#registro" className="hover:text-white transition-colors">Registro</a>
        </nav>

        <button
          onClick={() => navigate('/login')}
          className="flex items-center gap-2 bg-[#1a2235] hover:bg-[#232d45] border border-[#374151] text-white text-sm font-medium px-4 py-2 rounded-lg transition-colors"
        >
          <span aria-hidden>→</span>
          Iniciar Sesión
        </button>
      </header>

      {/* Hero */}
      <section
        className="relative overflow-hidden"
        style={{
          background:
            'radial-gradient(ellipse 80% 60% at 30% 20%, rgba(59,130,246,0.12) 0%, transparent 60%), radial-gradient(ellipse 60% 50% at 80% 80%, rgba(249,115,22,0.10) 0%, transparent 60%), linear-gradient(180deg, #0a0f1a 0%, #0d1526 100%)',
        }}
      >
        <div className="max-w-4xl mx-auto px-8 py-28 text-center">
          <span className="inline-block bg-[#f97316]/10 border border-[#f97316]/30 text-[#f97316] text-xs font-bold tracking-widest px-4 py-2 rounded-full mb-8">
            TRAZABILIDAD · CONTROL · VISIBILIDAD
          </span>

          <h1 className="text-4xl sm:text-5xl font-extrabold text-white leading-tight mb-6">
            ¿Sabes dónde están tus activos retornables en este momento?
          </h1>

          <p className="text-[#9ca3af] text-lg leading-relaxed mb-10 max-w-2xl mx-auto">
            Convierte contenedores, racks y embalajes retornables en activos inteligentes mediante
            trazabilidad, gobernanza y monitoreo en tiempo real.
          </p>

          <div className="flex flex-col sm:flex-row items-center justify-center gap-4">
            <button
              onClick={() => navigate('/login')}
              className="flex items-center gap-2 bg-white text-[#0a0f1a] font-semibold px-6 py-3 rounded-lg hover:bg-gray-100 transition-colors"
            >
              <Rocket size={18} />
              Solicitar acceso
            </button>
            <button
              onClick={() => navigate('/login')}
              className="flex items-center gap-2 bg-[#3b82f6] hover:bg-[#2563eb] text-white font-semibold px-6 py-3 rounded-lg transition-colors"
            >
              <Eye size={18} />
              Ver Plataforma
            </button>
          </div>
        </div>
      </section>
    </div>
  )
}
