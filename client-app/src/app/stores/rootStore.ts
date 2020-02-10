import RecipeStore from "./recipeStore";
import UserStore from "./userStore";
import { createContext } from "react";
import { configure } from "mobx";

configure({enforceActions: 'always'})

export class RootStore {
    recipeStore: RecipeStore;
    userStore: UserStore;

    constructor() {
        this.recipeStore = new RecipeStore(this);
        this.userStore = new UserStore(this);
    }
}

export const RootStoreContext = createContext(new RootStore());
