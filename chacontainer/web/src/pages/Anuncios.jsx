import { useState } from 'react'
import { Star, Phone, Mail, MapPin, Package, Wrench, ShoppingBag, Tag, Calendar, AlertCircle, Search } from 'lucide-react'
import { useApp } from '../context/AppContext'
import { ads } from '../data/mockData'

const TYPE_CONFIG = {
  venta: { label: 'Venta', bg: 'bg-emerald-900/40', text: 'text-emerald-400', border: 'border-emerald-700/50' },
  compra: { label: 'Compra', bg: 'bg-blue-900/40', text: 'text-blue-400', border: 'border-blue-700/50' },
  servicio: { label: 'Servicio', bg: 'bg-violet-900/40', text: 'text-violet-400', border: 'border-violet-700/50' },
}

const CATEGORY_ICON = {
  IBC: Package,
  Tambores: Package,
  Pallets: Package,
  Servicios: Wrench,
}

const CATEGORIES = ['Todas', 'IBC', 'Tambores', 'Pallets', 'Servicios']
const TYPES = ['Todos', 'venta', 'compra', 'servicio']

function AdCard({ ad, starred, onToggleStar }) {
  const [contactOpen, setContactOpen] = useState(false)
  const type = TYPE_CONFIG[ad.type]
  const Icon = CATEGORY_ICON[ad.category] || Package

  return (
    <div className={`bg-[#111827] border rounded-xl p-4 flex flex-col gap-3 transition-all ${
      starred ? 'border-[#f97316]/50' : 'border-[#1f2937]'
    }`}>
      {/* Header */}
      <div className="flex items-start justify-between gap-2">
        <div className="flex items-center gap-2 flex-wrap">
          <span className={`text-xs font-semibold px-2 py-0.5 rounded-full border ${type.bg} ${type.text} ${type.border}`}>
            {ad.type === 'venta' && <ShoppingBag size={10} className="inline mr-1" />}
            {ad.type === 'compra' && <Tag size={10} className="inline mr-1" />}
            {ad.type === 'servicio' && <Wrench size={10} className="inline mr-1" />}
            {type.label}
          </span>
          <span className="text-xs text-[#6b7280] font-mono bg-[#0a0f1a] px-2 py-0.5 rounded">{ad.category}</span>
          {ad.urgent && (
            <span className="text-xs font-semibold text-red-400 bg-red-900/30 border border-red-700/40 px-2 py-0.5 rounded-full flex items-center gap-1">
              <AlertCircle size={10} />
              Urgente
            </span>
          )}
        </div>
        <button
          onClick={() => onToggleStar(ad.id)}
          className={`shrink-0 p-1.5 rounded-lg transition-all ${
            starred
              ? 'text-[#f97316] bg-[#f97316]/10'
              : 'text-[#6b7280] hover:text-[#f97316] hover:bg-[#f97316]/10'
          }`}
          title={starred ? 'Quitar de destacados' : 'Destacar anuncio'}
        >
          <Star size={16} fill={starred ? 'currentColor' : 'none'} />
        </button>
      </div>

      {/* Title + Description */}
      <div>
        <div className="flex items-start gap-2 mb-1">
          <div className="w-8 h-8 bg-[#f97316]/10 rounded-lg flex items-center justify-center shrink-0 mt-0.5">
            <Icon size={16} className="text-[#f97316]" />
          </div>
          <h3 className="text-[#f9fafb] font-semibold text-sm leading-tight">{ad.title}</h3>
        </div>
        <p className="text-[#6b7280] text-xs leading-relaxed ml-10">{ad.description}</p>
      </div>

      {/* Details row */}
      <div className="flex flex-wrap gap-3 text-xs text-[#9ca3af] ml-10">
        {ad.quantity && (
          <span className="flex items-center gap-1">
            <Package size={11} className="text-[#6b7280]" />
            {ad.quantity} {ad.unit}
          </span>
        )}
        <span className="flex items-center gap-1">
          <MapPin size={11} className="text-[#6b7280]" />
          {ad.location}
        </span>
        <span className="flex items-center gap-1">
          <Calendar size={11} className="text-[#6b7280]" />
          {ad.postedAt}
        </span>
      </div>

      {/* Price + Contact */}
      <div className="flex items-center justify-between mt-auto">
        <div>
          {ad.price ? (
            <>
              <p className="text-[#f97316] text-lg font-bold font-mono">
                ${ad.price.toLocaleString('es-CL')}
              </p>
              <p className="text-[#6b7280] text-xs">por {ad.unit}</p>
            </>
          ) : (
            <p className="text-[#9ca3af] text-sm italic">Precio a convenir</p>
          )}
        </div>
        <button
          onClick={() => setContactOpen(o => !o)}
          className="text-xs font-medium px-3 py-1.5 rounded-lg border border-[#374151] text-[#9ca3af] hover:border-[#f97316]/40 hover:text-[#f97316] transition-colors"
        >
          {contactOpen ? 'Ocultar' : 'Contactar'}
        </button>
      </div>

      {/* Contact info panel */}
      {contactOpen && (
        <div className="bg-[#0a0f1a] border border-[#1f2937] rounded-lg p-3 space-y-1.5 text-xs">
          <p className="text-[#f9fafb] font-medium">{ad.company}</p>
          <a href={`mailto:${ad.contact}`} className="flex items-center gap-1.5 text-[#9ca3af] hover:text-[#f97316] transition-colors">
            <Mail size={11} />
            {ad.contact}
          </a>
          <a href={`tel:${ad.phone}`} className="flex items-center gap-1.5 text-[#9ca3af] hover:text-[#f97316] transition-colors">
            <Phone size={11} />
            {ad.phone}
          </a>
        </div>
      )}
    </div>
  )
}

export default function Anuncios() {
  const { starredAds, toggleStarAd } = useApp()
  const [activeCategory, setActiveCategory] = useState('Todas')
  const [activeType, setActiveType] = useState('Todos')
  const [onlyStarred, setOnlyStarred] = useState(false)
  const [search, setSearch] = useState('')

  const filtered = ads.filter(ad => {
    if (activeCategory !== 'Todas' && ad.category !== activeCategory) return false
    if (activeType !== 'Todos' && ad.type !== activeType) return false
    if (onlyStarred && !starredAds.has(ad.id)) return false
    if (search) {
      const q = search.toLowerCase()
      if (!ad.title.toLowerCase().includes(q) && !ad.description.toLowerCase().includes(q) && !ad.company.toLowerCase().includes(q)) return false
    }
    return true
  })

  const starCount = starredAds.size

  return (
    <div className="space-y-6">
      {/* Page header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-xl font-bold text-[#f9fafb]">Anuncios</h1>
          <p className="text-[#6b7280] text-sm">Clasificados de compra, venta y servicios de contenedores</p>
        </div>
        <button
          onClick={() => setOnlyStarred(o => !o)}
          className={`flex items-center gap-2 px-4 py-2 rounded-lg border text-sm font-medium transition-all ${
            onlyStarred
              ? 'bg-[#f97316]/10 border-[#f97316]/40 text-[#f97316]'
              : 'border-[#374151] text-[#9ca3af] hover:border-[#f97316]/30 hover:text-[#f97316]'
          }`}
        >
          <Star size={15} fill={onlyStarred ? 'currentColor' : 'none'} />
          Destacados
          {starCount > 0 && (
            <span className={`text-xs font-bold w-5 h-5 rounded-full flex items-center justify-center ${
              onlyStarred ? 'bg-[#f97316] text-white' : 'bg-[#374151] text-[#9ca3af]'
            }`}>
              {starCount}
            </span>
          )}
        </button>
      </div>

      <div className="flex flex-col sm:flex-row gap-4">
        {/* Sidebar filters */}
        <div className="w-full sm:w-48 shrink-0 space-y-4">
          {/* Search */}
          <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
            <p className="text-xs text-[#6b7280] font-semibold mb-3 uppercase tracking-wide">Buscar</p>
            <div className="relative">
              <Search size={13} className="absolute left-2.5 top-1/2 -translate-y-1/2 text-[#6b7280]" />
              <input
                type="text"
                value={search}
                onChange={e => setSearch(e.target.value)}
                placeholder="Buscar..."
                className="w-full bg-[#0a0f1a] border border-[#374151] rounded-lg pl-7 pr-3 py-2 text-xs text-[#f9fafb] placeholder-[#6b7280] focus:outline-none focus:border-[#f97316]/50"
              />
            </div>
          </div>

          {/* Category */}
          <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
            <p className="text-xs text-[#6b7280] font-semibold mb-3 uppercase tracking-wide">Categoría</p>
            <div className="space-y-1">
              {CATEGORIES.map(cat => (
                <button key={cat} onClick={() => setActiveCategory(cat)}
                  className={`w-full text-left px-3 py-2 rounded-lg text-sm transition-colors ${
                    activeCategory === cat
                      ? 'bg-[#f97316]/10 text-[#f97316] font-medium'
                      : 'text-[#9ca3af] hover:text-[#f9fafb] hover:bg-white/5'
                  }`}>
                  {cat}
                </button>
              ))}
            </div>
          </div>

          {/* Type */}
          <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
            <p className="text-xs text-[#6b7280] font-semibold mb-3 uppercase tracking-wide">Tipo</p>
            <div className="space-y-1">
              {TYPES.map(t => (
                <button key={t} onClick={() => setActiveType(t)}
                  className={`w-full text-left px-3 py-2 rounded-lg text-sm capitalize transition-colors ${
                    activeType === t
                      ? 'bg-[#f97316]/10 text-[#f97316] font-medium'
                      : 'text-[#9ca3af] hover:text-[#f9fafb] hover:bg-white/5'
                  }`}>
                  {t === 'Todos' ? 'Todos' : TYPE_CONFIG[t]?.label}
                </button>
              ))}
            </div>
          </div>
        </div>

        {/* Ad grid */}
        <div className="flex-1">
          {filtered.length === 0 ? (
            <div className="flex flex-col items-center justify-center py-20 text-[#6b7280]">
              <Star size={36} className="mb-3 opacity-20" />
              <p className="text-sm">No hay anuncios que coincidan con los filtros</p>
            </div>
          ) : (
            <div className="grid grid-cols-1 xl:grid-cols-2 gap-4">
              {filtered.map(ad => (
                <AdCard
                  key={ad.id}
                  ad={ad}
                  starred={starredAds.has(ad.id)}
                  onToggleStar={toggleStarAd}
                />
              ))}
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
