import { RouteComponentProps } from "react-router-dom";
import { Grid } from "semantic-ui-react";
import React from "react";
import { observer } from "mobx-react-lite";

interface DetailParams {
    id: string;
}

const RecipeForm : React.FC<RouteComponentProps<DetailParams>> = ({}) => {
    return (
        <Grid>
            
        </Grid>
    );
};

export default observer(RecipeForm);