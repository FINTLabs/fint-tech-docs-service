import { Component, OnInit, Input, ElementRef } from '@angular/core';
import { RepoModel } from '../api/RepoModel';
import { RepositoriesService } from "../api/repositories.service";

/**
 * Component responsible for displaying the details of
 * a repository. It expects to receive a `RepoModel` object
 * as input:
 *
 * ```html
 * <app-repo-card [repo]="repoModel"></app-repo-card>
 * ```
 */
@Component({
  selector: 'app-repo-card',
  templateUrl: './repo-card.component.html',
  styleUrls: ['./repo-card.component.scss']
})
export class RepoCardComponent implements OnInit {
  @Input() repo: RepoModel;

  get id() {
    return this.repoService.friendlyName(this.repo);
  }
  constructor(public element: ElementRef, private repoService: RepositoriesService) { }

  ngOnInit() {
  }

}
