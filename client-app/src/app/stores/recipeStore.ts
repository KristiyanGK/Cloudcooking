import { RootStore } from "./rootStore";
import { observable, action, runInAction, computed } from "mobx";
import { IRecipe } from "../models/recipe";
import agent from "../api/agent";
import { toast } from 'react-toastify';
import { history } from '../..';
import { SyntheticEvent } from "react";

export default class RecipeStore {
    rootStore: RootStore; 

    constructor(rootStore: RootStore) {
        this.rootStore = rootStore;
    }

    @observable recipeRegistry = new Map();
    @observable recipe: IRecipe | null = null;
    @observable loadingInitial = false;
    @observable submitting = false;
    @observable target = '';

    @computed get recipesByCategory() {
      let result = this.groupRecipesByCategory(Array.from(this.recipeRegistry.values()));
      return result;
    }

    groupRecipesByCategory(recipes: IRecipe[]) {
      const sortedRecipes = recipes.sort((a, b) => a.category.name.localeCompare(b.category.name));

      return Object.entries(sortedRecipes.reduce((recipes, recipe) => {
        const category = recipe.category;
        recipes[category.name] = recipes[category.name] ? [...recipes[category.name], recipe] : [recipe];
        return recipes;
      }, {} as {[key: string]: IRecipe[]}));
    }

    @action loadRecipes = async () => {
        this.loadingInitial = true;
        try {
          const recipes = await agent.Recipes.list(); 
          runInAction('loading recipes', () => {
            recipes.forEach(recipe => {
              this.recipeRegistry.set(recipe.id, recipe);
            });
            this.loadingInitial = false;
          })
        } catch (error) {
          runInAction('load recipes error', () => {
            this.loadingInitial = false;
          })
        }
    };

    @action loadRecipe = async (id: string) => {
        let recipe = this.getRecipe(id);
        if (recipe) {
          this.recipe = recipe;
          return recipe;
        } else {
          this.loadingInitial = true;
          try {
            recipe = await agent.Recipes.details(id);
            runInAction('getting recipe',() => {
              this.recipe = recipe;
              this.recipeRegistry.set(recipe.id, recipe);
              this.loadingInitial = false;
            })
            return recipe;
          } catch (error) {
            runInAction('get recipe error', () => {
              this.loadingInitial = false;
            })
            console.log(error);
          }
        }
    }

    @action clearRecipe = () => {
        this.recipe = null;
    }

    @action createRecipe = async (recipe: IRecipe) => {
        this.submitting = true;
        try {
          const recipeResult = await agent.Recipes.create(recipe);
          runInAction('create recipe', () => {
            this.recipeRegistry.set(recipeResult.id, recipeResult);
            this.submitting = false;
          })
          history.push(`/recipes/${recipeResult.id}`)
        } catch (error) {
          runInAction('create recipe error', () => {
            this.submitting = false;
          })
          toast.error('Problem submitting data');
          console.log(error.response);
        }
    };

    @action editRecipe = async (recipe: IRecipe) => {
        this.submitting = true;
        try {
          await agent.Recipes.update(recipe);
          runInAction('editing recipe', () => {
            this.recipeRegistry.set(recipe.id, recipe);
            this.recipe = recipe;
            this.submitting = false;
          })
          history.push(`/recipes/${recipe.id}`)
        } catch (error) {
          runInAction('edit recipe error', () => {
            this.submitting = false;
          })
          toast.error('Problem submitting data');
          console.log(error);
        }
    };

    @action deleteActivity = async (event: SyntheticEvent<HTMLButtonElement>, id: string) => {
        this.submitting = true;
        this.target = event.currentTarget.name;
        try {
          await agent.Recipes.delete(id);
          runInAction('deleting recipe', () => {
            this.recipeRegistry.delete(id);
            this.submitting = false;
            this.target = '';
          })
        } catch (error) {
          runInAction('delete activity error', () => {
            this.submitting = false;
            this.target = '';
          })
          console.log(error);
        }
    };

    getRecipe = (id: string) => {
        return this.recipeRegistry.get(id);
    }
}