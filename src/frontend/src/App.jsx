import { useCallback, useEffect, useState } from "react";
import api from "./api";
import debounce from "lodash.debounce";
import { LinePlot, createDataSet } from "./components/chart";

function App() {
  const [backendData, setBackendData] = useState(
    createDataSet([], [], "Analytical solution", "red")
  );
  const [numericalCentral, setNumericalCentral] = useState(
    createDataSet([], [], "Numerical-central difference scheme", "green")
  );
  const [numericalDir, setNumericalDir] = useState(
    createDataSet([], [], "Numerical-directional difference scheme", "blue")
  );

  const [errors, setErrors] = useState({
    central: 0,
    dir: 0
  });

  const [n, setN] = useState(9.0);
  const [eps, setEps] = useState(1);
  const [task, setTask] = useState("3");

  const setLines = async () => {
    var vals = { n, eps, task };
    if (n == 0) vals.n = 10;
    if (eps == 0) vals.eps = 1;
    var centralData = await api.getNumerical(vals.task, vals.eps, vals.n, "central");
    var dirData = await api.getNumerical(vals.task, vals.eps, vals.n, "directional");
    if (centralData.status == 200 && dirData.status == 200) {
      setNumericalCentral(createDataSet(centralData.data.xVals, centralData.data.yVals, "Numerical-central difference scheme", "green"));
      setNumericalDir(createDataSet(dirData.data.xVals, dirData.data.yVals, "Numerical-directional difference scheme", "blue"));
      setErrors({ central: centralData.data.err, dir: dirData.data.err })
    }
  };

  const analytical = async () => {
    var vals = { n, eps, task };
    if (vals.eps == 0) vals.eps = 1;
    var data = await api.getAnalytical(vals.task, vals.eps);
    if (data.status == 200) {
      setBackendData(
        createDataSet(data.data.xVals, data.data.yVals, "original", "red")
      );
    }
  };

  const setLinesCallBack = useCallback(setLines, []);
  const setAnalyticalCallBack = useCallback(analytical)

  useEffect(() => {
    setAnalyticalCallBack();
    setLinesCallBack();
    return () => { };
  }, []);

  const debounceN = debounce(query => {
    if (!query) return setN(0)
    var parsed = parseFloat(query)
    if (parsed <= 0) parsed = 1;
    if (parsed > 10000) parsed = 10000;
    setN(parsed)
  }, 10)

  const debounceEps = debounce(query => {
    if (!query) return setEps(0)
    var parsed = parseFloat(query)
    if (parsed > 1) parsed = 1;
    setEps(parsed)
  }, 10)

  useEffect(() => {
    setLines(n, eps, task)
  }, [n, eps, task]);

  useEffect(() => {
    analytical()
  }, [eps, task]);

  return (
    <div className="flex px-4">
      <div className="flex flex-col p-4">
        <h2 className="text-lg">ODE solver</h2>
        <div className="flex flex-col">
          <div className="flex flex-col my-6">
            <label className="text-gray-800-text-small mb-2">Grid size</label>
            <input
              type="number"
              className="outline-none ring hover:shadow-xl ring-green-400 rounded-lg px-4 mx-auto"
              value={n}
              onInput={(e) => debounceN(e.target.value)}
            />
          </div>
          <div className="flex flex-col my-6">
            <label className="text-gray-800-text-small mb-2">Epsilon</label>
            <input
              type="number"
              className="outline-none ring hover:shadow-xl ring-green-400 rounded-lg px-4 mx-auto"
              value={eps}
              onInput={(e) => debounceEps(e.target.value)}
            />
          </div>
          <div className="flex flex-col my-6">
            <label className="text-gray-800-text-small mb-2">Task</label>
            <select className="outline-none ring ring-green-400 rounded-lg hover:shadow-xl bg-white py-1 px-2" value={task} onChange={(e) => { setTask(e.target.value) }}>
              <option value="1">Task #1</option>
              <option value="2">Task #2</option>
              <option value="3">Task #3 </option>
            </select>
          </div>
          <div className="flex flex-col my-6 space-y-4">
            <p>Error for central scheme: {errors.central}</p>
            <p>Error for directional scheme: {errors.dir}</p>
          </div>
        </div>
      </div>
      <div className="flex flex-1">
        {LinePlot([backendData, numericalCentral, numericalDir])}
      </div>
    </div>
  );
}

export default App;
