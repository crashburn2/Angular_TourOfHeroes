import { Injectable } from '@angular/core';
import { Hero } from './hero';
import { HEROES } from './mock-heroes';
import { Observable, of } from 'rxjs';
import { MessageService } from './message.service';

@Injectable({
  providedIn: 'root'
})
export class HeroService {

  constructor(
    private messageService: MessageService) { }

  getHeroes(): Observable<Hero[]> {
    // TODO: send the message _after_ fetching the heroes
    this.messageService.add('HeroService: fetched heroes');
    return of(HEROES);
  }
  getHero(id: number | undefined): Observable<Hero | undefined> {
    // TODO: send the message _after_ fetching the hero
    if (id) {
    this.messageService.add(`HeroService: fetched hero id=${id}`);
    }
    return of(HEROES.find(hero => hero.id === id));
  }
}
//(method) Array<Hero>.find(predicate: (value: Hero, index: number, obj: Hero[]) => unknown, thisArg?: any): Hero | undefined