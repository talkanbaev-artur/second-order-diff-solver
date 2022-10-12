import { useEffect, useState } from "react";
import api from "./api";
import { LinePlot, createDataSet } from "./components/chart";


function App() {

  const [backendData, setBackendData] = useState(null);

  useEffect(() => {
    const func = async () => {
      var data = await api.getTypes();
      if (data.status == 200) {
        setBackendData(data.data);
      };
    }
    func()
    return () => { }
  }, []);

  var x = [];
  var y = [];
  for (let i = 0; i < 100; i++) {
    x[i] = i * 0.1;
    y[i] = x[i] ** Math.E;
  }

  var data = [createDataSet(x, y, "original", "rgb(255,10,10)")]
  return (
    <div className="">
      <div className="">
        {JSON.stringify(backendData)}
        {LinePlot(data)}
      </div>
    </div>
  )
}

export default App
