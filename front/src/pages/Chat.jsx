import React, { useEffect, useState } from "react";

import { useNavigate } from "react-router-dom";

import styled from "styled-components";

import { useCookies } from "react-cookie";
import { isExistCookie } from "../utils/CookieChecker";
import ChatContainer from "../components/ChatContainer";
import Contacts from "../components/Contacts";
import Welcome from "../components/Welcome";
import { useGlobalData } from "../context/context";
import { socketHost, host } from "../utils/APIRoutes";
import axios from "axios";

export default function Chat() {
  const [cookies] = useCookies(["auth"]);

  const storage = useGlobalData();
  const navigate = useNavigate();

  const [initLoading, setInitLoading] = useState(false);
  const [userId, setUserId] = useState("");

  useEffect(() => {
    if (!isExistCookie(cookies)) {
      navigate("/login");
    } else {
      setUserId(cookies.auth);
      setInitLoading(true);
      storage.setUserName(cookies.auth);
    }
  }, []);

  return (
    <>
      <Container>
        <div className="container">
          <Contacts userName={userId} />
          {!initLoading ? (
            <Welcome userName={userId} />
          ) : (
            <ChatContainer userName={userId} />
          )}
        </div>
      </Container>
    </>
  );
}

const Container = styled.div`
  height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 1rem;
  align-items: center;
  background-color: #131324;
  .container {
    height: 85vh;
    width: 85vw;
    background-color: #00000076;
    display: grid;
    grid-template-columns: 25% 75%;
    @media screen and (min-width: 720px) and (max-width: 1080px) {
      grid-template-columns: 35% 65%;
    }
  }
`;
