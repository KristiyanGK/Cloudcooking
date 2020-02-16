export interface IRecipe {
    id: string;
    title: string;
    description: string;
    usedProducts: string;
    cookingTime: Number;
    category: string;
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

    constructor(init?: IRecipeFormValues) {
        Object.assign(this, init);
    }
}