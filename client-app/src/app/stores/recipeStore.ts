import { RootStore } from "./rootStore";

export default class RecipeStore {
    rootStore: RootStore;

    constructor(rootStore: RootStore) {
        this.rootStore = rootStore;
    }
}