import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import { AppProvider } from './context/AppContext'
import Layout from './components/Layout'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import Marketplace from './pages/Marketplace'
import Projects from './pages/Projects'
import Traceability from './pages/Traceability'
import Calculator from './pages/Calculator'
import Backoffice from './pages/Backoffice'

export default function App() {
  return (
    <AppProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<Layout />}>
            <Route index element={<Navigate to="/dashboard" replace />} />
            <Route path="dashboard" element={<Dashboard />} />
            <Route path="marketplace" element={<Marketplace />} />
            <Route path="projects" element={<Projects />} />
            <Route path="traceability" element={<Traceability />} />
            <Route path="calculator" element={<Calculator />} />
            <Route path="backoffice" element={<Backoffice />} />
          </Route>
          <Route path="*" element={<Navigate to="/login" replace />} />
        </Routes>
      </BrowserRouter>
    </AppProvider>
  )
}
