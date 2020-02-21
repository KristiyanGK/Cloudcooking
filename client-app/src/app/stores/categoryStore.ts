import { RootStore } from "./rootStore";
import { observable, action, runInAction, computed, keys } from "mobx";
import agent from "../api/agent";
import { ICategory } from "../models/category";

export default class CategoryStore {
    rootStore: RootStore; 

    constructor(rootStore: RootStore) {
        this.rootStore = rootStore;
    }

    @observable categoryRegistry = new Map<string, ICategory>();
    @observable loadingInitial = false;
    @observable submitting = false;

    @computed get categoriesAsOptions() {
        let result = new Array();
        this.categoryRegistry.forEach((val : ICategory, key: string) => {
            result.push({key: key, text: val.name, value: val.id});
        });

        return result;
    }

    @action loadCategories = async () => {
        this.loadingInitial = true;
        try {
          const categories = await agent.Categories.list();
          runInAction('loading categories', () => {
            categories.forEach(category => {
              this.categoryRegistry.set(category.id, category);
            });
            this.loadingInitial = false;
          })
        } catch (error) {
          runInAction('load categories error', () => {
            this.loadingInitial = false;
          })
        }
    };
}
