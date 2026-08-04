import { useState } from 'react'
import { ShoppingCart, X, Plus, Minus, Package, CheckCircle, Truck, ChevronDown } from 'lucide-react'
import { useApp } from '../context/AppContext'
import { products } from '../data/mockData'

const categories = ['Todos', 'Contenedor Colapsable', 'Contenedor / Caja Rígida', 'KLT', 'Rack Metálico', 'Tarima / Pallet', 'Charola', 'Separador', 'Servicios']
const warehouses = [
  { id: 'bodega_edomex', label: 'Bodega Estado de México' },
  { id: 'bodega_queretaro', label: 'Bodega Querétaro' },
  { id: 'bodega_puebla', label: 'Bodega Puebla' },
]

function StockBadge({ stock }) {
  if (stock >= 50) return <span className="text-emerald-400 text-xs font-mono">{stock} un</span>
  if (stock >= 10) return <span className="text-amber-400 text-xs font-mono">{stock} un</span>
  if (stock > 0) return <span className="text-red-400 text-xs font-mono">{stock} un ⚠</span>
  return <span className="text-[#6b7280] text-xs font-mono">Sin stock</span>
}

function CartDrawer({ onClose }) {
  const { cart, removeFromCart, updateCartQty, cartTotal, clearCart, selectedWarehouse } = useApp()
  const [checkedOut, setCheckedOut] = useState(false)
  const [deliveryDate, setDeliveryDate] = useState('')
  const [showModal, setShowModal] = useState(false)

  if (checkedOut) {
    return (
      <div className="fixed right-0 top-0 h-full w-96 bg-[#111827] border-l border-[#1f2937] z-50 flex flex-col items-center justify-center p-8">
        <CheckCircle size={48} className="text-emerald-400 mb-4" />
        <h3 className="text-xl font-bold text-[#f9fafb] mb-2">¡Pedido Confirmado!</h3>
        <p className="text-[#6b7280] text-sm text-center mb-2">ORD-{Math.floor(Math.random() * 1000 + 2850)}</p>
        <p className="text-[#9ca3af] text-xs text-center mb-6">Recibirá confirmación por email. Tiempo estimado: 24-48 hrs.</p>
        <button onClick={() => { setCheckedOut(false); clearCart(); onClose() }}
          className="bg-[#f97316] text-white px-6 py-2 rounded-lg text-sm font-medium">
          Volver al Marketplace
        </button>
      </div>
    )
  }

  return (
    <>
      <div className="fixed inset-0 bg-black/50 z-40" onClick={onClose} />
      <div className="fixed right-0 top-0 h-full w-96 bg-[#111827] border-l border-[#1f2937] z-50 flex flex-col">
        <div className="p-4 border-b border-[#1f2937] flex items-center justify-between">
          <h3 className="font-bold text-[#f9fafb] flex items-center gap-2">
            <ShoppingCart size={18} className="text-[#f97316]" />
            Carrito de Compras
          </h3>
          <button onClick={onClose} className="text-[#6b7280] hover:text-[#f9fafb]">
            <X size={18} />
          </button>
        </div>

        {cart.length === 0 ? (
          <div className="flex-1 flex flex-col items-center justify-center text-[#6b7280]">
            <ShoppingCart size={40} className="mb-3 opacity-30" />
            <p className="text-sm">Carrito vacío</p>
          </div>
        ) : (
          <>
            <div className="flex-1 overflow-y-auto p-4 space-y-3">
              {cart.map(item => (
                <div key={item.id} className="bg-[#1a2235] border border-[#1f2937] rounded-lg p-3">
                  <div className="flex items-start justify-between mb-2">
                    <div className="flex-1 min-w-0">
                      <p className="text-[#f9fafb] text-xs font-medium leading-tight">{item.name}</p>
                      <p className="text-[#6b7280] text-xs font-mono">{item.sku}</p>
                    </div>
                    <button onClick={() => removeFromCart(item.id)} className="text-[#6b7280] hover:text-red-400 ml-2">
                      <X size={14} />
                    </button>
                  </div>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center gap-2">
                      <button onClick={() => updateCartQty(item.id, item.qty - 1)}
                        className="w-6 h-6 bg-[#111827] border border-[#374151] rounded flex items-center justify-center text-[#9ca3af] hover:text-white">
                        <Minus size={10} />
                      </button>
                      <span className="text-[#f9fafb] text-sm font-mono w-6 text-center">{item.qty}</span>
                      <button onClick={() => updateCartQty(item.id, item.qty + 1)}
                        className="w-6 h-6 bg-[#111827] border border-[#374151] rounded flex items-center justify-center text-[#9ca3af] hover:text-white">
                        <Plus size={10} />
                      </button>
                    </div>
                    <p className="text-[#f97316] text-sm font-mono font-bold">
                      ${(item.price * item.qty).toLocaleString('es-CL')}
                    </p>
                  </div>
                </div>
              ))}
            </div>

            <div className="p-4 border-t border-[#1f2937] space-y-3">
              <div className="flex justify-between text-sm">
                <span className="text-[#9ca3af]">Subtotal</span>
                <span className="text-[#f9fafb] font-mono font-bold">${cartTotal.toLocaleString('es-CL')}</span>
              </div>
              <div className="flex justify-between text-xs">
                <span className="text-[#6b7280]">IVA (19%)</span>
                <span className="text-[#9ca3af] font-mono">${Math.round(cartTotal * 0.19).toLocaleString('es-CL')}</span>
              </div>
              <div className="flex justify-between text-sm font-bold border-t border-[#1f2937] pt-2">
                <span className="text-[#f9fafb]">Total</span>
                <span className="text-[#f97316] font-mono">${Math.round(cartTotal * 1.19).toLocaleString('es-CL')}</span>
              </div>
              <button onClick={() => setShowModal(true)}
                className="w-full bg-[#f97316] hover:bg-[#ea6c0a] text-white py-2.5 rounded-lg text-sm font-semibold transition-colors">
                Confirmar Pedido
              </button>
            </div>
          </>
        )}
      </div>

      {showModal && (
        <div className="fixed inset-0 bg-black/70 z-[60] flex items-center justify-center p-4">
          <div className="bg-[#111827] border border-[#1f2937] rounded-xl w-full max-w-md p-6">
            <h3 className="font-bold text-[#f9fafb] text-lg mb-4 flex items-center gap-2">
              <Truck size={20} className="text-[#f97316]" />
              Confirmar Pedido
            </h3>
            <div className="space-y-4">
              <div>
                <label className="text-xs text-[#9ca3af] block mb-1">Bodega de Despacho</label>
                <div className="bg-[#0a0f1a] border border-[#374151] rounded-lg px-3 py-2 text-[#f9fafb] text-sm">
                  {warehouses.find(w => w.id === selectedWarehouse)?.label}
                </div>
              </div>
              <div>
                <label className="text-xs text-[#9ca3af] block mb-1">Fecha de Entrega Deseada</label>
                <input type="date" value={deliveryDate} onChange={e => setDeliveryDate(e.target.value)}
                  className="w-full bg-[#0a0f1a] border border-[#374151] rounded-lg px-3 py-2 text-[#f9fafb] text-sm focus:outline-none focus:border-[#f97316]" />
              </div>
              <div className="bg-[#0a0f1a] rounded-lg p-3 space-y-1">
                {cart.map(item => (
                  <div key={item.id} className="flex justify-between text-xs">
                    <span className="text-[#9ca3af]">{item.qty}x {item.name}</span>
                    <span className="text-[#f9fafb] font-mono">${(item.price * item.qty).toLocaleString('es-CL')}</span>
                  </div>
                ))}
                <div className="border-t border-[#1f2937] pt-1 mt-1 flex justify-between text-sm font-bold">
                  <span className="text-[#f9fafb]">Total c/IVA</span>
                  <span className="text-[#f97316] font-mono">${Math.round(cartTotal * 1.19).toLocaleString('es-CL')}</span>
                </div>
              </div>
              <div className="flex gap-3">
                <button onClick={() => setShowModal(false)}
                  className="flex-1 border border-[#374151] text-[#9ca3af] py-2 rounded-lg text-sm hover:border-[#6b7280]">
                  Cancelar
                </button>
                <button onClick={() => { setShowModal(false); setCheckedOut(true) }}
                  className="flex-1 bg-[#f97316] text-white py-2 rounded-lg text-sm font-semibold hover:bg-[#ea6c0a]">
                  Confirmar
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </>
  )
}

export default function Marketplace() {
  const { addToCart, setCartOpen, cartOpen, selectedWarehouse, setSelectedWarehouse, cartCount } = useApp()
  const [activeCategory, setActiveCategory] = useState('Todos')
  const [quantities, setQuantities] = useState({})

  const filtered = activeCategory === 'Todos' ? products : products.filter(p => p.category === activeCategory)

  function getQty(id) { return quantities[id] || 1 }
  function setQty(id, val) { setQuantities(prev => ({ ...prev, [id]: Math.max(1, val) })) }

  function handleAdd(product) {
    addToCart(product, getQty(product.id))
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-xl font-bold text-[#f9fafb]">Marketplace</h1>
          <p className="text-[#6b7280] text-sm">Catálogo de contenedores y servicios reutilizables</p>
        </div>
        <button onClick={() => setCartOpen(true)}
          className="relative flex items-center gap-2 bg-[#f97316] hover:bg-[#ea6c0a] text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors">
          <ShoppingCart size={16} />
          <span>Carrito</span>
          {cartCount > 0 && (
            <span className="bg-white text-[#f97316] text-xs font-bold w-5 h-5 rounded-full flex items-center justify-center">
              {cartCount}
            </span>
          )}
        </button>
      </div>

      <div className="flex flex-col sm:flex-row gap-4">
        {/* Filters sidebar */}
        <div className="w-full sm:w-48 shrink-0 space-y-4">
          <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
            <p className="text-xs text-[#6b7280] font-semibold mb-3 uppercase tracking-wide">Categoría</p>
            <div className="space-y-1">
              {categories.map(cat => (
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

          <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-4">
            <p className="text-xs text-[#6b7280] font-semibold mb-3 uppercase tracking-wide">Bodega</p>
            <div className="space-y-1">
              {warehouses.map(w => (
                <button key={w.id} onClick={() => setSelectedWarehouse(w.id)}
                  className={`w-full text-left px-3 py-2 rounded-lg text-xs transition-colors ${
                    selectedWarehouse === w.id
                      ? 'bg-blue-500/10 text-blue-400 font-medium'
                      : 'text-[#9ca3af] hover:text-[#f9fafb] hover:bg-white/5'
                  }`}>
                  {w.label}
                </button>
              ))}
            </div>
          </div>
        </div>

        {/* Product grid */}
        <div className="flex-1 grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 gap-4">
          {filtered.map(product => {
            const stockHere = product.stock[selectedWarehouse]
            const isService = product.category === 'Servicios'
            return (
              <div key={product.id} className="bg-[#111827] border border-[#1f2937] rounded-xl p-4 flex flex-col gap-3">
                <div className="flex items-start justify-between">
                  <div className="w-10 h-10 bg-[#f97316]/10 rounded-lg flex items-center justify-center">
                    <Package size={20} className="text-[#f97316]" />
                  </div>
                  <span className="text-xs font-mono text-[#6b7280] bg-[#0a0f1a] px-2 py-0.5 rounded">{product.category}</span>
                </div>

                <div>
                  <h3 className="text-[#f9fafb] font-semibold text-sm leading-tight">{product.name}</h3>
                  <p className="text-[#6b7280] text-xs font-mono mt-0.5">{product.sku}</p>
                </div>

                <div className="flex items-center justify-between">
                  <div>
                    <p className="text-[#f97316] text-lg font-bold font-mono">
                      ${product.price.toLocaleString('es-CL')}
                    </p>
                    <p className="text-[#6b7280] text-xs">por {product.unit}</p>
                  </div>
                  <div className="text-right">
                    <p className="text-xs text-[#6b7280] mb-0.5">Stock aquí</p>
                    <StockBadge stock={isService ? '∞' : stockHere} />
                  </div>
                </div>

                {product.co2Saved > 0 && (
                  <div className="flex items-center gap-1.5 bg-emerald-900/20 border border-emerald-900/40 rounded-lg px-2.5 py-1.5">
                    <div className="w-1.5 h-1.5 bg-emerald-500 rounded-full" />
                    <span className="text-emerald-400 text-xs">CO₂ evitado: {product.co2Saved} kg/unidad</span>
                  </div>
                )}

                <div className="flex items-center gap-2 mt-auto">
                  <div className="flex items-center gap-1 border border-[#374151] rounded-lg">
                    <button onClick={() => setQty(product.id, getQty(product.id) - 1)}
                      className="px-2 py-1 text-[#9ca3af] hover:text-white">
                      <Minus size={12} />
                    </button>
                    <span className="text-[#f9fafb] text-sm font-mono w-8 text-center">{getQty(product.id)}</span>
                    <button onClick={() => setQty(product.id, getQty(product.id) + 1)}
                      className="px-2 py-1 text-[#9ca3af] hover:text-white">
                      <Plus size={12} />
                    </button>
                  </div>
                  <button onClick={() => handleAdd(product)}
                    disabled={!isService && stockHere === 0}
                    className="flex-1 bg-[#f97316]/10 hover:bg-[#f97316]/20 border border-[#f97316]/30 text-[#f97316] text-xs font-medium py-2 rounded-lg transition-colors disabled:opacity-40 disabled:cursor-not-allowed">
                    Agregar
                  </button>
                </div>
              </div>
            )
          })}
        </div>
      </div>

      {cartOpen && <CartDrawer onClose={() => setCartOpen(false)} />}
    </div>
  )
}
