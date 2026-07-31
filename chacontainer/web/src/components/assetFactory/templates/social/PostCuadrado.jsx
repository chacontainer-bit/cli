import { forwardRef } from 'react'
import BrandMark from '../BrandMark'
import { BRAND } from '../../../../lib/assetFactory/brand'

const PostCuadrado = forwardRef(function PostCuadrado({ data }, ref) {
  const { titulo, subtitulo, imagen } = data
  return (
    <div
      ref={ref}
      style={{
        width: 1080,
        height: 1080,
        position: 'relative',
        overflow: 'hidden',
        background: BRAND.surfaceDark,
        fontFamily: BRAND.fontFamily,
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
      }}
    >
      {imagen && (
        <>
          <img src={imagen} alt="" style={{ position: 'absolute', inset: 0, width: '100%', height: '100%', objectFit: 'cover' }} />
          <div
            style={{
              position: 'absolute',
              inset: 0,
              background: `linear-gradient(180deg, ${BRAND.surfaceDark}55 0%, ${BRAND.surfaceDark}ee 68%, ${BRAND.surfaceDark} 100%)`,
            }}
          />
        </>
      )}

      <div style={{ position: 'relative', padding: '64px 72px 0' }}>
        <BrandMark size={30} textSize={20} />
      </div>

      <div style={{ position: 'relative', padding: '0 72px 88px' }}>
        <div style={{ width: 84, height: 8, background: BRAND.color, borderRadius: 4, marginBottom: 36 }} />
        <h1
          style={{
            color: BRAND.textLight,
            fontSize: 68,
            fontWeight: 800,
            lineHeight: 1.08,
            margin: 0,
            maxWidth: 880,
          }}
        >
          {titulo || 'Título de la pieza'}
        </h1>
        {subtitulo && (
          <p style={{ color: BRAND.textMuted, fontSize: 30, lineHeight: 1.4, marginTop: 28, maxWidth: 820 }}>
            {subtitulo}
          </p>
        )}
      </div>
    </div>
  )
})

export default PostCuadrado
