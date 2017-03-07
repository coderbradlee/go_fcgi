#ifndef PDF_HPP
#define PDF_HPP
#include <stdbool.h>
#include <stdio.h>
#include "../wkhtmltox/include/pdf.h"

void Test(int n);
/* Print out loading progress information */
void progress_changed(wkhtmltopdf_converter * c, int p);

/* Print loading phase information */
void phase_changed(wkhtmltopdf_converter * c);
/* Print a message to stderr when an error occurs */
void error(wkhtmltopdf_converter * c, const char * msg);
void warning(wkhtmltopdf_converter * c, const char * msg);
void Convert();
#endif