import React, { useContext } from "react";

class GlobalData {
  socket;
  userName;

  setUserName = (name) => {
    this.userName = name;
  };

  setSocket = (socket) => {
    this.socket = socket;
  };
}

export const GlobalDataContext = React.createContext(new GlobalData());

export const useGlobalData = () => {
  return useContext(GlobalDataContext);
};

export default GlobalData;
