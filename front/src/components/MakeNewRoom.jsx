import axios from "axios";
import React from "react";
import styled from "styled-components";
import { host } from "../utils/APIRoutes";

export default function MakeNewRoom({ setChattingList }) {
  const generateRandomRoomName = () => {
    const adjectives = ["Happy", "Sunny", "Cheerful", "Bright", "Colorful"];
    const nouns = ["Chat", "Room", "Conversation", "Exchange", "Dialogue"];
    const randomAdjective =
      adjectives[Math.floor(Math.random() * adjectives.length)];
    const randomNoun = nouns[Math.floor(Math.random() * nouns.length)];
    return `${randomAdjective} ${randomNoun}`;
  };

  const handleClick = async () => {
    const name = generateRandomRoomName();
    alert(name + " will be created");
    const res = await axios.post(host + "/make-room", {
      Name: name,
    });

    if (res.status !== 200) {
      alert("Failed to create new room");
    } else {
      const list = await axios.get(host + "/room-list");
      setChattingList(list.data.result);
    }
  };

  return (
    <Button onClick={handleClick}>
      <span>New Room</span>
    </Button>
  );
}

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
