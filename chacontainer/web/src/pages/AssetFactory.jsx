import { useMemo, useRef, useState } from 'react'
import {
  Layers, ArrowLeft, Download, Image as ImageIcon, Handshake, Presentation, Rocket, ChevronRight, Loader2,
} from 'lucide-react'
import { CATEGORIES, EXPORT_KIND, findAssetType } from '../lib/assetFactory/catalog'
import { buildFileName } from '../lib/assetFactory/fileName'
import { exportNodeAsPng, exportNodesAsPdf, exportHtmlString } from '../lib/assetFactory/exporters'
import { getTemplate } from '../components/assetFactory/templates'
import { emptyFormValue } from '../lib/assetFactory/formValue'
import AssetForm from '../components/assetFactory/AssetForm'
import PreviewFrame from '../components/assetFactory/PreviewFrame'

const CATEGORY_ICON = {
  'redes-sociales': ImageIcon,
  comercial: Handshake,
  presentaciones: Presentation,
  landing: Rocket,
}

function CategoryGrid({ onSelect }) {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-4">
      {CATEGORIES.map(cat => {
        const Icon = CATEGORY_ICON[cat.id] || Layers
        return (
          <button
            key={cat.id}
            onClick={() => onSelect(cat.id)}
            className="bg-[#111827] border border-[#1f2937] rounded-xl p-5 text-left hover:border-[#f97316]/50 hover:bg-[#f97316]/5 transition-all"
          >
            <div className="w-10 h-10 rounded-lg bg-[#f97316]/10 flex items-center justify-center mb-4">
              <Icon size={20} className="text-[#f97316]" />
            </div>
            <h3 className="text-[#f9fafb] font-semibold text-sm">{cat.label}</h3>
            <p className="text-[#6b7280] text-xs mt-1.5">{cat.description}</p>
            <p className="text-[#6b7280] text-xs mt-3">{cat.types.length} opciones</p>
          </button>
        )
      })}
    </div>
  )
}

function TypeGrid({ category, onSelect, onBack }) {
  return (
    <div className="space-y-4">
      <button onClick={onBack} className="flex items-center gap-1.5 text-xs text-[#9ca3af] hover:text-[#f9fafb]">
        <ArrowLeft size={14} /> Categorías
      </button>
      <div>
        <h2 className="text-[#f9fafb] font-semibold text-lg">{category.label}</h2>
        <p className="text-[#6b7280] text-sm mt-0.5">{category.description}</p>
      </div>
      <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
        {category.types.map(type => (
          <button
            key={type.id}
            onClick={() => onSelect(type.id)}
            className="bg-[#111827] border border-[#1f2937] rounded-xl p-4 text-left hover:border-[#f97316]/50 hover:bg-[#f97316]/5 transition-all flex items-center justify-between"
          >
            <div>
              <h3 className="text-[#f9fafb] font-semibold text-sm">{type.label}</h3>
              <p className="text-[#6b7280] text-xs mt-1">{type.description}</p>
            </div>
            <ChevronRight size={16} className="text-[#6b7280] flex-shrink-0 ml-3" />
          </button>
        ))}
      </div>
    </div>
  )
}

// Templates 'multi' (propuesta comercial, deck) renderizan varias páginas de
// tamaño real a través de registerPage(index, node). Ese mismo nodo DOM se usa
// para el preview (escalado visualmente con transform) y para el export a
// PDF (offsetWidth/offsetHeight no cambian por el transform), así preview y
// archivo descargado nunca se desincronizan.
function MultiPageScaledPreview({ Component, data, registerPage, pageWidth, maxWidth = 300 }) {
  const scale = Math.min(1, maxWidth / pageWidth)
  return (
    <div className="flex gap-4 overflow-x-auto w-full pb-2">
      <Component data={data} previewScale={scale} registerPage={registerPage} />
    </div>
  )
}

function Editor({ categoryId, typeId, onBack, onChangeType }) {
  const { category, type: assetType } = findAssetType(categoryId, typeId)
  const [formValue, setFormValue] = useState(emptyFormValue)
  const [exporting, setExporting] = useState(false)
  const singleRef = useRef(null)
  const pageRefs = useRef([])
  const template = getTemplate(typeId)

  function registerPage(index, node) {
    pageRefs.current[index] = node
  }

  const fileExtension = useMemo(() => {
    if (assetType.exportKind === EXPORT_KIND.PNG) return 'png'
    if (assetType.exportKind === EXPORT_KIND.PDF) return 'pdf'
    return 'html'
  }, [assetType])

  async function handleDownload() {
    setExporting(true)
    try {
      const fileName = buildFileName({
        categoryId,
        typeId,
        titulo: formValue.titulo,
        extension: fileExtension,
      })

      if (assetType.exportKind === EXPORT_KIND.PNG) {
        await exportNodeAsPng(singleRef.current, fileName)
      } else if (assetType.exportKind === EXPORT_KIND.PDF) {
        const nodes = template.kind === 'multi' ? pageRefs.current.filter(Boolean) : [singleRef.current]
        await exportNodesAsPdf(nodes, fileName, assetType.pageSize)
      } else if (assetType.exportKind === EXPORT_KIND.HTML) {
        const html = template.buildHtml(formValue)
        exportHtmlString(html, fileName)
      }
    } finally {
      setExporting(false)
    }
  }

  return (
    <div className="space-y-4">
      <div className="flex items-center gap-2 text-xs text-[#9ca3af]">
        <button onClick={onBack} className="flex items-center gap-1.5 hover:text-[#f9fafb]">
          <ArrowLeft size={14} /> Categorías
        </button>
        <span>/</span>
        <button onClick={onChangeType} className="hover:text-[#f9fafb]">{category.label}</button>
        <span>/</span>
        <span className="text-[#f9fafb]">{assetType.label}</span>
      </div>

      <div className="grid grid-cols-1 xl:grid-cols-[380px_1fr] gap-6">
        <div className="bg-[#111827] border border-[#1f2937] rounded-xl p-5 space-y-5">
          <div>
            <h2 className="text-[#f9fafb] font-semibold text-sm">{assetType.label}</h2>
            <p className="text-[#6b7280] text-xs mt-1">{assetType.description}</p>
          </div>
          <AssetForm assetType={assetType} value={formValue} onChange={setFormValue} />
          <button
            onClick={handleDownload}
            disabled={exporting}
            className="w-full flex items-center justify-center gap-2 bg-[#f97316] hover:bg-[#ea6c0a] disabled:opacity-60 text-white font-semibold text-sm py-3 rounded-lg transition-colors"
          >
            {exporting ? <Loader2 size={16} className="animate-spin" /> : <Download size={16} />}
            {exporting ? 'Generando…' : `Descargar .${fileExtension.toUpperCase()}`}
          </button>
        </div>

        <div className="bg-[#0a0f1a] border border-[#1f2937] rounded-xl p-6 flex flex-col items-center">
          <p className="text-[#6b7280] text-xs mb-4 self-start">Vista previa en vivo</p>
          <div className="w-full flex items-start justify-center overflow-auto">
            {template?.kind === 'single' && (
              <PreviewFrame
                width={assetType.width || assetType.pageSize?.[0]}
                height={assetType.height || assetType.pageSize?.[1]}
              >
                <template.Component data={formValue} ref={singleRef} />
              </PreviewFrame>
            )}
            {template?.kind === 'multi' && (
              <MultiPageScaledPreview
                Component={template.Component}
                data={formValue}
                registerPage={registerPage}
                pageWidth={assetType.pageSize[0]}
              />
            )}
            {template?.kind === 'html' && <template.Component data={formValue} />}
            {!template && <p className="text-[#6b7280] text-sm">Template no disponible.</p>}
          </div>
        </div>
      </div>
    </div>
  )
}

export default function AssetFactory() {
  const [categoryId, setCategoryId] = useState(null)
  const [typeId, setTypeId] = useState(null)

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-xl font-bold text-[#f9fafb] flex items-center gap-2">
          <Layers size={20} className="text-[#f97316]" />
          Fábrica de Assets de Marca
        </h1>
        <p className="text-[#6b7280] text-sm mt-1">
          Elige una categoría, completa el contenido y descarga la pieza con el diseño real de{' '}
          <span className="text-[#f97316]">CHACONTAINER</span>.
        </p>
      </div>

      {!categoryId && <CategoryGrid onSelect={setCategoryId} />}

      {categoryId && !typeId && (
        <TypeGrid
          category={CATEGORIES.find(c => c.id === categoryId)}
          onSelect={setTypeId}
          onBack={() => setCategoryId(null)}
        />
      )}

      {categoryId && typeId && (
        <Editor
          categoryId={categoryId}
          typeId={typeId}
          onBack={() => {
            setCategoryId(null)
            setTypeId(null)
          }}
          onChangeType={() => setTypeId(null)}
        />
      )}
    </div>
  )
}
