import { Component, OnInit, ViewChildren, QueryList } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { RepositoriesService } from '../api/repositories.service';
import { RepoModel } from '../api/RepoModel';
import { AppComponent } from '../app.component';
import { RepoCardComponent } from "../repo-card/repo-card.component";

/**
 * Component responsible for reading in the list of repositories,
 * and handling the filtering of repositories according to type.
 */
@Component({
  selector: 'app-repo-list',
  templateUrl: './repo-list.component.html',
  styleUrls: ['./repo-list.component.scss']
})
export class RepoListComponent implements OnInit {
  @ViewChildren(RepoCardComponent) cards: QueryList<RepoCardComponent>;
  repositories: RepoModel[] = [];
  type: string;

  constructor(private parent: AppComponent, private repo: RepositoriesService, private route: ActivatedRoute) { }

  ngOnInit() {
    this.route.params.subscribe((param: any) => {
      this.type = param && param.type ? param.type : null;
      this.repo.all().subscribe(result => this.reposReceived(<RepoModel[]>result, this.type));
      if (param.class) {
        setTimeout(() => {
          const selectedCard = this.cards.find(c => c.id === param.class);
          if (selectedCard) {
            selectedCard.element.nativeElement.scrollIntoView({block: 'start', behavior: 'smooth'});
          }
        });
      }
    });

    this.parent.searchChanged.subscribe(searchValue => {
      this.repo.all().subscribe(result => {
        this.reposReceived(<RepoModel[]>result, this.type);
        this.repositories = this.repositories.filter(model => {
          return Object.keys(model).some(key => {
            const val = model[key];
            if (typeof val === 'string') {
              return val.toLowerCase().indexOf(searchValue.toLowerCase()) > -1;
            }
            return false;
          });
        });
      });
    });
  }

  reposReceived(result: RepoModel[], filter?: string) {
    this.repositories = (filter ? result.filter(m => m.lang === filter) : result);
    // this.repositories = result;
  }
}