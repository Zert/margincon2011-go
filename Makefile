all:
	pdflatex margin_go.tex

clean:
	${RM} *.{aux,log,nav,out,pdf,snm,toc,vrb} *~
