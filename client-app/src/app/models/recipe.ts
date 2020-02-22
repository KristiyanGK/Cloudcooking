import { ICategory } from "./category";

export interface IRecipeEnvelope {
    recipes: IRecipe[];
    count: number;
}

export interface IRecipe {
    id: string;
    title: string;
    description: string;
    usedProducts: string;
    cookingTime: Number;
    category: ICategory;
    categoryId: string;
    user?: string;
    createdAt: Date;
}

export interface IRecipeFormValues extends Partial<IRecipe> {

}

export class RecipeFormValues implements IRecipeFormValues {
    id?: string = undefined;
    title: string = '';
    category?: ICategory = undefined;
    description: string = '';
    usedProducts: string = '';
    cookingTime: Number = 0;
    user?: string = '';

    constructor(init?: IRecipeFormValues) {
        Object.assign(this, init);
    }
}