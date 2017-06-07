import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AppComponent } from './app.component';
import { RepoListComponent } from "./repo-list/repo-list.component";

const routes: Routes = [
  { path: '', component: RepoListComponent },
  { path: ':type', component: RepoListComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
