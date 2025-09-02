import Header from './components/Header';
import SidebarLeft from './components/SidebarLeft';
import SidebarRight from './components/SidebarRight';
import Content from './components/Content';
import Footer from './components/Footer';
import './App.css';

export default function App() {
  return (
    <div className="container">
      <Header />
      <div className="main">
        <SidebarLeft />
        <Content />
        <SidebarRight />
      </div>
      <Footer />
    </div>
  );
}