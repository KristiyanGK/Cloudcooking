export interface IRecipe {
    id: string;
    title: string;
    description: string;
    usedProducts: string;
    cookingTime: Number;
    category: string;
    categoryId: string;
    user?: string;
}

export interface IRecipeFormValues extends Partial<IRecipe> {

}

export class RecipeFormValues implements IRecipeFormValues {
    id?: string = undefined;
    title: string = '';
    category: string = '';
    description: string = '';
    usedProducts: string = '';
    cookingTime: Number = 0;
    user?: string = '';

    constructor(init?: IRecipeFormValues) {
        Object.assign(this, init);
    }
}