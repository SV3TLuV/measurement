import {persistor, store} from "./lib/store";
import {Provider} from "react-redux";
import {PersistGate} from "redux-persist/integration/react";
import {ConfigProvider} from "antd";
import ru_RU from 'antd/locale/ru_RU';
import {RouterProvider} from "react-router-dom";
import {router} from "./lib/router/router.tsx";

function App() {
    return (
      <Provider store={store}>
          <PersistGate loading={null} persistor={persistor}>
              <ConfigProvider locale={ru_RU}>
                  <RouterProvider router={router}/>
              </ConfigProvider>
          </PersistGate>
      </Provider>
    )
}

export default App
