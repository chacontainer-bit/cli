import { BRAND } from './brand'

function esc(str) {
  return String(str || '').replace(/[&<>"']/g, c => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' }[c]))
}

// Genera un documento HTML autocontenido (CSS inline, sin dependencias
// externas obligatorias) con el diseño real de marca, listo para subir a
// cualquier hosting como landing de campaña.
export function buildLandingHtml(data) {
  const { titulo, subtitulo, cta, imagen, datosClave } = data
  const benefits = (datosClave || []).filter(d => d.label || d.value)
  const title = titulo || 'Título de la campaña'

  const benefitCards = benefits
    .map(
      b => `
        <div class="benefit">
          <div class="benefit-dot"></div>
          <h3>${esc(b.label)}</h3>
          <p>${esc(b.value)}</p>
        </div>`
    )
    .join('')

  return `<!doctype html>
<html lang="es">
<head>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width, initial-scale=1" />
<title>${esc(title)} · ${BRAND.name}</title>
<style>
  :root {
    --brand: ${BRAND.color};
    --brand-dark: ${BRAND.colorDark};
    --surface: ${BRAND.surface};
    --surface-dark: ${BRAND.surfaceDark};
    --surface-card: ${BRAND.surfaceCard};
    --border: ${BRAND.border};
    --text: ${BRAND.textLight};
    --text-muted: ${BRAND.textMuted};
    --text-faint: ${BRAND.textFaint};
  }
  * { box-sizing: border-box; }
  body {
    margin: 0;
    background: var(--surface-dark);
    color: var(--text);
    font-family: 'Inter', system-ui, -apple-system, sans-serif;
    -webkit-font-smoothing: antialiased;
  }
  .brand { display: flex; align-items: center; gap: 10px; }
  .brand-dot { width: 12px; height: 12px; border-radius: 3px; background: var(--brand); }
  .brand span { color: var(--brand); font-weight: 800; letter-spacing: 0.14em; font-size: 15px; }
  header { padding: 28px 48px; }
  .hero {
    position: relative;
    padding: 64px 48px 96px;
    overflow: hidden;
    text-align: center;
  }
  .hero img.bg {
    position: absolute; inset: 0; width: 100%; height: 100%; object-fit: cover; opacity: 0.28; z-index: 0;
  }
  .hero .overlay {
    position: absolute; inset: 0; background: linear-gradient(180deg, transparent 0%, var(--surface-dark) 90%); z-index: 1;
  }
  .hero-content { position: relative; z-index: 2; max-width: 760px; margin: 0 auto; }
  .accent-bar { width: 72px; height: 6px; background: var(--brand); border-radius: 3px; margin: 0 auto 28px; }
  h1 { font-size: clamp(32px, 6vw, 56px); font-weight: 800; line-height: 1.1; margin: 0; }
  .subtitle { color: var(--text-muted); font-size: 19px; line-height: 1.6; margin: 22px auto 0; max-width: 640px; }
  .cta {
    display: inline-block; margin-top: 36px; background: var(--brand); color: #fff; font-weight: 700;
    font-size: 17px; padding: 16px 36px; border-radius: 12px; text-decoration: none; transition: background 0.15s;
  }
  .cta:hover { background: var(--brand-dark); }
  .benefits { padding: 16px 48px 88px; max-width: 1040px; margin: 0 auto; }
  .benefits-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(240px, 1fr)); gap: 20px; }
  .benefit {
    background: var(--surface-card); border: 1px solid var(--border); border-radius: 16px; padding: 26px;
  }
  .benefit-dot { width: 10px; height: 10px; border-radius: 50%; background: var(--brand); margin-bottom: 16px; }
  .benefit h3 { font-size: 17px; margin: 0 0 8px; }
  .benefit p { color: var(--text-muted); font-size: 14px; line-height: 1.55; margin: 0; }
  footer {
    border-top: 1px solid var(--border); padding: 28px 48px; text-align: center;
    color: var(--text-faint); font-size: 13px;
  }
</style>
</head>
<body>
  <header>
    <div class="brand"><span class="brand-dot" style="display:inline-block;width:10px;height:10px;border-radius:3px;background:var(--brand);"></span><span>${BRAND.name}</span></div>
  </header>

  <section class="hero">
    ${imagen ? `<img class="bg" src="${imagen}" alt="" /><div class="overlay"></div>` : ''}
    <div class="hero-content">
      <div class="accent-bar"></div>
      <h1>${esc(title)}</h1>
      ${subtitulo ? `<p class="subtitle">${esc(subtitulo)}</p>` : ''}
      ${cta ? `<a class="cta" href="#contacto">${esc(cta)}</a>` : ''}
    </div>
  </section>

  ${benefits.length > 0 ? `
  <section class="benefits">
    <div class="benefits-grid">${benefitCards}</div>
  </section>` : ''}

  <footer>${BRAND.name} · ${BRAND.tagline}</footer>
</body>
</html>`
}
