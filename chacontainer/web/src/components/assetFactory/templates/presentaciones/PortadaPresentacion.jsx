import { forwardRef } from 'react'
import BrandMark from '../BrandMark'
import { BRAND, SLIDE_16_9 } from '../../../../lib/assetFactory/brand'

const PortadaPresentacion = forwardRef(function PortadaPresentacion({ data }, ref) {
  const { titulo, subtitulo, imagen } = data
  return (
    <div
      ref={ref}
      style={{
        width: SLIDE_16_9.width,
        height: SLIDE_16_9.height,
        background: BRAND.surfaceDark,
        fontFamily: BRAND.fontFamily,
        boxSizing: 'border-box',
        position: 'relative',
        overflow: 'hidden',
        padding: '56px 80px',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
      }}
    >
      {imagen && (
        <>
          <img src={imagen} alt="" style={{ position: 'absolute', inset: 0, width: '100%', height: '100%', objectFit: 'cover', opacity: 0.3 }} />
          <div style={{ position: 'absolute', inset: 0, background: `linear-gradient(135deg, ${BRAND.surfaceDark}f5 0%, ${BRAND.surfaceDark}aa 100%)` }} />
        </>
      )}
      <div style={{ position: 'relative' }}>
        <BrandMark size={28} textSize={18} />
      </div>
      <div style={{ position: 'relative' }}>
        <div style={{ width: 84, height: 8, background: BRAND.color, borderRadius: 4, marginBottom: 28 }} />
        <h1 style={{ color: BRAND.textLight, fontSize: 62, fontWeight: 800, lineHeight: 1.08, margin: 0, maxWidth: 960 }}>
          {titulo || 'Título de la presentación'}
        </h1>
        {subtitulo && <p style={{ color: BRAND.textMuted, fontSize: 26, marginTop: 22, maxWidth: 780 }}>{subtitulo}</p>}
      </div>
      <p style={{ position: 'relative', color: BRAND.textFaint, fontSize: 14, margin: 0 }}>
        {new Date().toLocaleDateString('es-MX', { day: 'numeric', month: 'long', year: 'numeric' })}
      </p>
    </div>
  )
})

export default PortadaPresentacion
