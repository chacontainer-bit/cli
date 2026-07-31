import { forwardRef } from 'react'
import BrandMark from '../BrandMark'
import { BRAND } from '../../../../lib/assetFactory/brand'

const BannerHorizontal = forwardRef(function BannerHorizontal({ data }, ref) {
  const { titulo, subtitulo, cta, imagen } = data
  return (
    <div
      ref={ref}
      style={{
        width: 1200,
        height: 628,
        background: BRAND.surfaceDark,
        fontFamily: BRAND.fontFamily,
        display: 'flex',
        boxSizing: 'border-box',
        overflow: 'hidden',
      }}
    >
      <div style={{ flex: imagen ? '0 0 56%' : '1 1 100%', padding: '64px 64px', display: 'flex', flexDirection: 'column', justifyContent: 'space-between' }}>
        <BrandMark size={26} textSize={17} />
        <div>
          <div style={{ width: 64, height: 6, background: BRAND.color, borderRadius: 3, marginBottom: 24 }} />
          <h1 style={{ color: BRAND.textLight, fontSize: 50, fontWeight: 800, lineHeight: 1.1, margin: 0 }}>
            {titulo || 'Título del banner'}
          </h1>
          {subtitulo && (
            <p style={{ color: BRAND.textMuted, fontSize: 22, lineHeight: 1.4, marginTop: 16, maxWidth: 480 }}>
              {subtitulo}
            </p>
          )}
        </div>
        {cta && (
          <div
            style={{
              display: 'inline-flex',
              alignSelf: 'flex-start',
              background: BRAND.color,
              color: '#ffffff',
              fontWeight: 700,
              fontSize: 18,
              padding: '14px 28px',
              borderRadius: 10,
            }}
          >
            {cta}
          </div>
        )}
      </div>

      {imagen && (
        <div style={{ flex: '0 0 44%', position: 'relative' }}>
          <img src={imagen} alt="" style={{ width: '100%', height: '100%', objectFit: 'cover' }} />
          <div
            style={{
              position: 'absolute',
              inset: 0,
              background: `linear-gradient(90deg, ${BRAND.surfaceDark} 0%, transparent 18%)`,
            }}
          />
        </div>
      )}
    </div>
  )
})

export default BannerHorizontal
