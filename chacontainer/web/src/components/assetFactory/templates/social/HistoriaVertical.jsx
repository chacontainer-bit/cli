import { forwardRef } from 'react'
import BrandMark from '../BrandMark'
import { BRAND } from '../../../../lib/assetFactory/brand'

const HistoriaVertical = forwardRef(function HistoriaVertical({ data }, ref) {
  const { titulo, subtitulo, cta, imagen, datosClave } = data
  const chips = (datosClave || []).filter(d => d.label || d.value)

  return (
    <div
      ref={ref}
      style={{
        width: 1080,
        height: 1920,
        position: 'relative',
        overflow: 'hidden',
        background: BRAND.surfaceDark,
        fontFamily: BRAND.fontFamily,
        display: 'flex',
        flexDirection: 'column',
        boxSizing: 'border-box',
      }}
    >
      {imagen && (
        <>
          <img src={imagen} alt="" style={{ position: 'absolute', inset: 0, width: '100%', height: '100%', objectFit: 'cover' }} />
          <div
            style={{
              position: 'absolute',
              inset: 0,
              background: `linear-gradient(180deg, ${BRAND.surfaceDark}88 0%, ${BRAND.surfaceDark}cc 55%, ${BRAND.surfaceDark} 100%)`,
            }}
          />
        </>
      )}

      <div style={{ position: 'relative', padding: '84px 64px 0' }}>
        <BrandMark size={32} textSize={21} />
      </div>

      <div style={{ position: 'relative', marginTop: 'auto', padding: '0 64px 120px' }}>
        <div style={{ width: 90, height: 8, background: BRAND.color, borderRadius: 4, marginBottom: 32 }} />
        <h1 style={{ color: BRAND.textLight, fontSize: 66, fontWeight: 800, lineHeight: 1.12, margin: 0 }}>
          {titulo || 'Título de la historia'}
        </h1>
        {subtitulo && (
          <p style={{ color: BRAND.textMuted, fontSize: 28, lineHeight: 1.4, marginTop: 24 }}>{subtitulo}</p>
        )}

        {chips.length > 0 && (
          <div style={{ display: 'flex', flexWrap: 'wrap', gap: 12, marginTop: 36 }}>
            {chips.map((c, i) => (
              <div
                key={i}
                style={{
                  background: `${BRAND.surfaceCard}`,
                  border: `1px solid ${BRAND.border}`,
                  borderRadius: 12,
                  padding: '14px 20px',
                }}
              >
                {c.label && <p style={{ color: BRAND.color, fontSize: 15, fontWeight: 700, margin: 0 }}>{c.label}</p>}
                {c.value && <p style={{ color: BRAND.textLight, fontSize: 18, margin: '4px 0 0' }}>{c.value}</p>}
              </div>
            ))}
          </div>
        )}

        {cta && (
          <div
            style={{
              display: 'inline-flex',
              background: BRAND.color,
              color: '#ffffff',
              fontWeight: 700,
              fontSize: 24,
              padding: '20px 40px',
              borderRadius: 14,
              marginTop: 44,
            }}
          >
            {cta}
          </div>
        )}
      </div>
    </div>
  )
})

export default HistoriaVertical
