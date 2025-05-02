
export default function Home() {
  return (
    <div className="h-screen flex items-center justify-center">
        <div className="bg-white p-6 rounded shadow">
            <h1 className="text-2xl">LEGO SETS API</h1>
            <p>Visit <code>/api/lego</code> to get a list of LEGO sets.</p>
            <p><span className="font-medium">API DOCS:</span> <span id="url"></span></p>
        </div>
    </div>
  );
}
