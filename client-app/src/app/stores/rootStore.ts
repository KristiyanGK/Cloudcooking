import RecipeStore from "./recipeStore";
import UserStore from "./userStore";
import { createContext } from "react";
import { configure } from "mobx";
import CommonStore from "./commonStore";
import ModalStore from "./modalStore";
import CategoryStore from "./categoryStore";
import ChatStore from "./chatStore";
import CommentStore from "./commentStore";

configure({enforceActions: 'always'})

export class RootStore {
    recipeStore: RecipeStore;
    userStore: UserStore;
    commonStore: CommonStore;
    modalStore: ModalStore;
    categoryStore: CategoryStore;
    chatStore: ChatStore;
    commentStore: CommentStore;

    constructor() {
        this.recipeStore = new RecipeStore(this);
        this.userStore = new UserStore(this);
        this.commonStore = new CommonStore(this);
        this.modalStore = new ModalStore(this);
        this.categoryStore = new CategoryStore(this);
        this.chatStore = new ChatStore(this);
        this.commentStore = new CommentStore(this);
    }
}

export const RootStoreContext = createContext(new RootStore());
