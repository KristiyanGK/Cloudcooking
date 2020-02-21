import React from 'react';
import { Segment, Grid, Icon } from 'semantic-ui-react';
import { IRecipe } from '../../../app/models/recipe';

const RecipeDetailedInfo: React.FC<{recipe: IRecipe}> = ({recipe}) => {
  return (
    <Segment.Group>   
      <Segment attached>
        <Grid verticalAlign='middle'>
          <Grid.Column width={1}>
            <Icon name='clock' size='large' color='teal' />
          </Grid.Column>
          <Grid.Column width={11}>
          {recipe.cookingTime}
          </Grid.Column>
        </Grid>
      </Segment>
      <Segment attached>
        <Grid verticalAlign='middle'>
          <Grid.Column width={1}>
            <Icon name='ordered list' size='large' color='teal' />
          </Grid.Column>
          <Grid.Column width={11}>
            {recipe.usedProducts}
          </Grid.Column>
        </Grid>
      </Segment>
      <Segment attached='top'>
        <Grid>
          <Grid.Column width={1}>
            <Icon size='large' color='teal' name='info' />
          </Grid.Column>
          <Grid.Column width={15}>
            <p>{recipe.description}</p>
          </Grid.Column>
        </Grid>
      </Segment>
    </Segment.Group>
  );
};

export default RecipeDetailedInfo;
