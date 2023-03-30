import { TestBed } from '@angular/core/testing';
import { AppComponent } from './app.component';
import { RouterModule } from '@angular/router';

describe('AppComponent', () => {
    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [AppComponent],
            imports: [ RouterModule ]
        }).compileComponents();
    });

    it('should create the app', () => {
        const fixture = TestBed.createComponent(AppComponent);
        const app = fixture.componentInstance;
        (expect as any)(app).toBeTruthy();
    });

    it(`should have as title 'Karaokifyy'`, () => {
        const fixture = TestBed.createComponent(AppComponent);
        const app = fixture.componentInstance;
        (expect as any)(app.getTitle()).toEqual('Karaokifyy');
    });
});