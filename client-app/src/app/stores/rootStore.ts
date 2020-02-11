import RecipeStore from "./recipeStore";
import UserStore from "./userStore";
import { createContext } from "react";
import { configure } from "mobx";
import CommonStore from "./commonStore";

configure({enforceActions: 'always'})

export class RootStore {
    recipeStore: RecipeStore;
    userStore: UserStore;
    commonStore: CommonStore;

    constructor() {
        this.recipeStore = new RecipeStore(this);
        this.userStore = new UserStore(this);
        this.commonStore = new CommonStore(this);
    }
}

export const RootStoreContext = createContext(new RootStore());
