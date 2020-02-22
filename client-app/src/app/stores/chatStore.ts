import { RootStore } from "./rootStore";
import { observable, action } from "mobx";
import { IMessage } from "../models/message";

export default class ChatStore {
    rootStore: RootStore;
    
    constructor(rootStore: RootStore) {
        this.rootStore = rootStore;
    }

    @observable messages: IMessage[] = new Array();
    ws: WebSocket | null = null

    @action pushMsg = (message: IMessage) => {
        this.messages.push(message)
    }

    openConnection = () => {
        this.ws = new WebSocket("ws://localhost:8080/api/chat?token=" + this.rootStore?.commonStore.token);

        this.ws.onmessage = (event) => {
            this.pushMsg(JSON.parse(event.data))
        }
    }

    addMessage = (message: IMessage) => {
        this.ws!.send(JSON.stringify(message))
    }
}