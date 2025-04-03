import React, { useEffect, useState } from "react";
import styled from "styled-components";
import ChatInput from "./ChatInput";
import Logout from "./Logout";
import MakeNewRoom from "./MakeNewRoom";
import { host, socketHost } from "../utils/APIRoutes";
import axios from "axios";
import { useGlobalData } from "../context/context";

export default function ChatContainer({ userName }) {
  const [chattingList, setChattingList] = useState([]);
  const [room, setRoom] = useState(false);

  const [chatContents, setChatContents] = useState([]);
  const [socketH, setSocketH] = useState(false);

  const storage = useGlobalData();

  useEffect(() => {
    if (!room) {
      async function fetchRoomList() {
        const list = await axios.get(host + "/room-list");
        setChattingList(list.data.result);
      }

      fetchRoomList();
      setChatContents([]);
      storage.setSocket(false);
    } else {
      console.log(room);
      async function fetchBeforeChat() {
        const res = await axios.get(host + `/enter-room?name=${room}`);
        if (chatContents.length === 0) {
          setChatContents(res.data.result);
        }
      }

      fetchBeforeChat();

      const socket = new WebSocket(socketHost);
      setSocketH(socket);
    }
  }, [room]);

  if (socketH) {
    socketH.onopen = () => {
      //webSocket이 맺어지고 난 후, 실행
      console.log(socketH.current.readyState);
      socketH.current.send("success");
    };

    socketH.onmessage = function (e) {
      const receiveData = JSON.parse(e.data);

      console.log(chatContents);

      if (chatContents.length === 0) {
        setChatContents([receiveData]);
      } else {
        setChatContents([...chatContents, receiveData]);
      }
    };

    socketH.onclose = function (e) {
      console.log(e);
      // alert("서버가 닫혀있기 떄문에 로그아웃 됩니다.");

      // document.cookie =
      //   "auth" +
      //   "=" +
      //   ("/" ? ";path=" + "/" : "") +
      //   ";expires=Thu, 01 Jan 1970 00:00:01 GMT";

      // window.location.replace("/login");
    };
  }

  const handleSendMsg = async (msg) => {
    if (!room) {
      alert("먼저 방에 입장해 주세요");
    } else {
      if (!socketH) {
        alert("socket 설정이 아직 진행되지 않았습니다.");
      } else {
        console.log(socketH);
        socketH.send(JSON.stringify({ Message: msg, Room: room }));
      }
    }
  };

  const enterRoom = async (name) => {
    setRoom(name);
  };

  const goBack = () => {
    setRoom(false);
  };

  return (
    <Container>
      <div className="chat-header">
        {!room ? (
          <div className="user-details">
            <div className="username">
              <h3>채팅방 리스트</h3>
            </div>
            <MakeNewRoom setChattingList={setChattingList} />
          </div>
        ) : (
          <Button onClick={() => goBack()}>뒤로 가기</Button>
        )}

        <Logout />
      </div>
      <div className="room-list">
        {!room ? (
          chattingList.length !== 0 ? (
            chattingList.map((result, index) => (
              <div
                className={`room`}
                key={index}
                onClick={() => {
                  enterRoom(result.name);
                }}
              >
                {result.name}
              </div>
            ))
          ) : (
            <div>ttest</div>
          )
        ) : (
          <div>
            {chatContents.length !== 0
              ? chatContents.map((result, index) => {
                  const isMyText = result.name === userName;
                  return (
                    <div
                      className={`message ${isMyText ? "sended" : "received"}`}
                      key={index}
                    >
                      <div className="content-box">
                        <div className="content-header">
                          <span>{result.name}</span>
                        </div>
                        <div className="content">{result.message}</div>
                      </div>
                    </div>
                  );
                })
              : null}
          </div>
        )}
      </div>
      {room && <ChatInput handleSendMsg={handleSendMsg} />}
    </Container>
  );
}

const Container = styled.div`
  display: grid;
  grid-template-rows: 10% 80% 10%;
  gap: 0.1rem;
  overflow: hidden;
  @media screen and (min-width: 720px) and (max-width: 1080px) {
    grid-template-rows: 15% 70% 15%;
  }
  .chat-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 2rem;
    .user-details {
      display: flex;
      align-items: center;
      gap: 1rem;
      .avatar {
        img {
          height: 3rem;
        }
      }
      .username {
        h3 {
          color: white;
        }
      }
    }
  }
  .room-list {
    padding: 1rem 2rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    overflow: auto;

    .room {
      padding: 1rem 1rem;
      background-color: yellow;
      cursor: hover;
    }

    &::-webkit-scrollbar {
      width: 0.2rem;
      &-thumb {
        background-color: #ffffff39;
        width: 0.1rem;
        border-radius: 1rem;
      }
    }
    .message {
      display: flex;
      align-items: center;

      .content-box {
        flex-direction: column;
        max-width: 40%;

        .content-header {
          display: flex;
          flex-direction: row;

          span {
            display: flex;
            color: #fff;
            align-items: flex-end;
            padding-left: 5px;
          }
          img {
            width: 50px;
            height: 50px;
            border-radius: 50%;
          }
        }
        .content {
          overflow-wrap: break-word;
          padding: 1rem;
          font-size: 1.1rem;
          border-radius: 1rem;
          color: #d1d1d1;
          @media screen and (min-width: 720px) and (max-width: 1080px) {
            max-width: 70%;
          }
        }
      }
    }
    .sended {
      justify-content: flex-end;
      .content {
        background-color: #4f04ff21;
      }
    }
    .recieved {
      justify-content: flex-start;
      .content {
        background-color: #9900ff20;
      }
    }
  }
`;

const Button = styled.button`
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0.5rem;
  border-radius: 0.5rem;
  background-color: #9a86f3;
  border: none;
  cursor: pointer;
  svg {
    font-size: 1.3rem;
    color: #ebe7ff;
  }
`;
