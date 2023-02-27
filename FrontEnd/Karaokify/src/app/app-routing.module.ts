import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { SearchScreenItIsComponent } from "./search-screen-it-is/search-screen-it-is.component";
import { HomeScreenComponent } from "./home-screen/home-screen.component";

const routes: Routes = [
  { path: '', component:HomeScreenComponent},
  { path: 'screen-search', component:SearchScreenItIsComponent}
];


@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]

})
export class AppRoutingModule { }
export const routingComponents = [SearchScreenItIsComponent]