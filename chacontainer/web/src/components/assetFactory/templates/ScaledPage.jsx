// Usado por templates 'multi' (varias páginas/slides) para poder mostrar una
// vista previa escalada SIN aplicar el transform al nodo que se captura para
// el PDF. El nodo con `ref` siempre queda a tamaño real: el scale vive en un
// wrapper intermedio, así offsetWidth/offsetHeight y la captura de
// html-to-image nunca se ven afectados por el zoom del preview.
export default function ScaledPage({ width, height, scale = 1, registerRef, style, children }) {
  if (scale >= 1) {
    return (
      <div ref={registerRef} style={{ ...style, width, height }}>
        {children}
      </div>
    )
  }

  return (
    <div style={{ width: width * scale, height: height * scale, overflow: 'hidden', flexShrink: 0 }}>
      <div style={{ width, height, transform: `scale(${scale})`, transformOrigin: 'top left' }}>
        <div ref={registerRef} style={{ ...style, width, height }}>
          {children}
        </div>
      </div>
    </div>
  )
}
